package srvregister

import (
	"context"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteRegistry(ctx context.Context, in *pbservices.DeleteRegistryRequest) (*emptypb.Empty, error) {
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
	reg, err := retrieveRegistryByUser(ctx, db, server.GetUserFromAccessToken(auth, in.GetAccessToken()))
	if err != nil {
		return nil, err
	}
	authCloser()

	return s.deleteRegistry(ctx, db, reg)
}

func (s *Server) deleteRegistry(ctx context.Context, db *sqlx.DB, reg *ent.Registry) (*emptypb.Empty, error) {
	err := storeql.DeleteFromDB(ctx, db, reg)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "DeleteFromDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}
