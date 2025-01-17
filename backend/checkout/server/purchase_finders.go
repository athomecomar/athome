package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func FindPurchase(ctx context.Context, db *sqlx.DB, oId, uId uint64) (*purchase.Purchase, error) {
	order, err := purchase.FindPurchaseUserScoped(ctx, db, oId, uId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "FindPurchase with order_id: %v", oId)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindPurchase: %v", err)
	}
	return order, nil
}

func FindLatestPurchase(ctx context.Context, db *sqlx.DB, uId uint64) (*purchase.Purchase, error) {
	order, err := purchase.FindLatestPurchase(ctx, db, uId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(xerrors.NotFound, "purchase.FindLatestPurchase")
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindLatestPurchase: %v", err)
	}
	return order, nil
}
