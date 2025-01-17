package server

import (
	"context"

	"github.com/athomecomar/athome/backend/messenger/ent"
	"github.com/athomecomar/athome/backend/messenger/messengerconf"
	"github.com/athomecomar/athome/pb/pbmessenger"
	"github.com/athomecomar/athome/pb/pbnotifier"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) Create(ctx context.Context, in *pbmessenger.CreateRequest) (*pbmessenger.CreateResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	userId, err := pbutil.GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	err = authCloser()
	if err != nil {
		return nil, err
	}
	noti, notiCloser, err := pbutil.ConnNotifier(ctx)
	if err != nil {
		return nil, err
	}
	defer notiCloser()
	return s.create(ctx, db, noti, in, userId)
}

func (s *Server) create(ctx context.Context, db *sqlx.DB, noti pbnotifier.NotificationsClient, in *pbmessenger.CreateRequest, uid uint64) (*pbmessenger.CreateResponse, error) {
	pbCreateMsg := in.GetMessage()
	msg := ent.NewMessage(uid, pbCreateMsg.GetReceiverId(), pbCreateMsg.GetBody())
	err := storeql.InsertIntoDB(ctx, db, msg)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}

	_, err = pbutil.CreateNotification(
		ctx,
		msg,
		pbCreateMsg.GetReceiverId(),
		messengerconf.GetNOTIFICATION_JWT_SECRET,
		messengerconf.GetNOTIFICATION_BODY(),
		pbutil.High,
		noti,
	)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "msg.ToPb: %v", err)
	}

	pbMsg, err := msg.ToPb()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "msg.ToPb: %v", err)
	}
	return &pbmessenger.CreateResponse{
			MessageId: msg.Id,
			Message:   pbMsg,
		},
		nil
}
