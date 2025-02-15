package srvviewer

import (
	"context"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) SearchAvailableShippings(ctx context.Context, in *pbservices.SearchAvailableShippingsRequest) (*pbservices.SearchAvailableShippingsResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}

	sem, semCloser, err := pbutil.ConnSemanticServiceProviders(ctx)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	users, usersCloser, err := pbutil.ConnUsersViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer usersCloser()

	return s.searchAvailableShippings(ctx, db, sem, users, in)
}

func (s *Server) searchAvailableShippings(ctx context.Context, db *sqlx.DB,
	sem pbsemantic.ServiceProvidersClient, users pbusers.ViewerClient,
	in *pbservices.SearchAvailableShippingsRequest) (*pbservices.SearchAvailableShippingsResponse, error) {
	dow, err := ent.DayOfWeekByName(in.GetDow())
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "DayOfWeekByName: %v", err)
	}

	cat, err := sem.RetrieveShippingCategories(ctx, &pbsemantic.RetrieveShippingCategoriesRequest{MaxVolWeight: in.GetMaxVolWeight()})
	if err != nil {
		return nil, err
	}
	var catIds []uint64
	for id := range cat.Categories {
		catIds = append(catIds, id)
	}

	candidates, err := ent.AvailableServicesInAnyCategory(ctx, db, dow, in.GetStart(), in.GetEnd(), catIds...)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "DayOfWeekByName: %v", err)
	}
	startInMinutes := in.GetStart().GetMinute() + in.GetStart().GetHour()*60
	endInMinutes := in.GetEnd().GetMinute() + in.GetEnd().GetHour()*60
	maxDuration := float64(endInMinutes - startInMinutes)

	var svcs []*ent.Service
	for _, svc := range candidates {
		realDuration := float64(svc.DurationInMinutes) * in.GetDistanceInKilometers()
		if realDuration <= maxDuration {
			svcs = append(svcs, svc)
		}
	}

	resp := &pbservices.SearchAvailableShippingsResponse{}
	resp.Services = make(map[uint64]*pbservices.Service)
	for _, svc := range svcs {
		resp.Services[svc.Id] = svc.ToPb()
	}
	return resp, nil
}
