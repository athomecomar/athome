package srvpurchases

import (
	"context"
	"math"

	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/ent/shipping"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveShippingMethods(ctx context.Context, in *pbcheckout.RetrieveShippingMethodsRequest) (*pbcheckout.RetrieveShippingMethodsResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := pbutil.GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	err = authCloser()
	if err != nil {
		return nil, err
	}
	o, err := server.FindLatestPurchase(ctx, db, uid)
	if err != nil {
		return nil, err
	}
	err = server.MustPrevState(ctx, db, o, sm.PurchaseShippingMethodSelected, uid)
	if err != nil {
		return nil, err
	}

	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()

	svcs, svcsCloser, err := pbutil.ConnServicesViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer svcsCloser()

	return s.retrieveShippingMethods(ctx, db,
		prods, svcs,
		in, o,
	)
}

func (s *Server) retrieveShippingMethods(
	ctx context.Context,
	db *sqlx.DB,

	prods pbproducts.ViewerClient,
	svcs pbservices.ViewerClient,

	in *pbcheckout.RetrieveShippingMethodsRequest,
	p *purchase.Purchase,

) (*pbcheckout.RetrieveShippingMethodsResponse, error) {
	start, err := pbutil.RestTimeOfDay(in.GetTime(), dispatchDelayMinutes)
	if err != nil {
		return nil, err
	}

	maxVolWeight, err := p.MaxVolWeight(ctx, prods)
	if err != nil {
		return nil, err
	}

	services, err := svcs.SearchAvailableShippings(ctx, &pbservices.SearchAvailableShippingsRequest{
		MaxVolWeight:         maxVolWeight,
		DistanceInKilometers: p.DistanceInKilometers,

		Dow:   in.GetDow(),
		Start: start,
		End:   in.GetTime(),
	})
	if err != nil {
		return nil, err
	}

	resp := &pbcheckout.RetrieveShippingMethodsResponse{}
	resp.ShippingMethods = make(map[uint64]*pbcheckout.ShippingMethod)
	for id, svc := range services.GetServices() {
		ppkm, err := shipping.CalculateShippingPricePerKilometer(ctx, db, svc.GetUserId(), svc.GetPrice())
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "CalculateShippingPricePerKilometer: %v", err)
		}
		price := ppkm.Float64() * p.DistanceInKilometers
		resp.ShippingMethods[id] = &pbcheckout.ShippingMethod{
			Amount:            price,
			DurationInMinutes: uint64(math.Ceil(p.DistanceInKilometers * float64(svc.DurationInMinutes))),
			UserId:            svc.GetUserId(),
			Title:             svc.GetTitle(),
		}
	}
	return resp, nil
}
