package order

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Order interface {
	storeql.Storable
	StateMachine() *sm.StateMachine
	OrderClass() class
}

func StateChanges(ctx context.Context, db *sqlx.DB, o Order) (scs []StateChange, err error) {
	switch o.OrderClass() {
	case Purchases:
		scs, err = stateChangesPurchases(ctx, db, o.GetId())
	}
	return
}

func stateChangesPurchases(ctx context.Context, db *sqlx.DB, oid uint64) (scs []StateChange, err error) {
	rows, err := storeql.WhereMany(ctx, db, &PurchaseStateChange{}, "order_id=$1", oid)
	if err != nil {
		err = errors.Wrap(err, "storeql.WhereMany")
		return
	}
	defer rows.Close()
	for rows.Next() {
		sc := &PurchaseStateChange{}
		err = rows.StructScan(sc)
		if err != nil {
			err = errors.Wrap(err, "StructScan")
			return
		}
		scs = append(scs, sc)
	}
	return
}