package ent

import (
	"context"
	"time"

	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbshared"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/currency"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Service struct {
	Id         uint64 `json:"id,omitempty"`
	UserId     uint64 `json:"user_id,omitempty"`
	AddressId  uint64 `json:"address_id,omitempty"`
	CalendarId uint64 `json:"calendar_id,omitempty"`

	Title             string       `json:"title,omitempty"`
	DurationInMinutes uint64       `json:"duration_in_minutes,omitempty"`
	PriceMin          currency.ARS `json:"price_min,omitempty"`
	PriceMax          currency.ARS `json:"price_max,omitempty"`
}

func FindService(ctx context.Context, db *sqlx.DB, id uint64) (*Service, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM services WHERE id=$1`, id)
	svc := &Service{}
	err := row.StructScan(svc)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return svc, nil
}

func (s *Service) User(ctx context.Context, user pbusers.ViewerClient) (*pbusers.User, error) {
	return user.RetrieveUser(ctx, &pbusers.RetrieveUserRequest{UserId: s.UserId})
}

func (s *Service) Calendar(ctx context.Context, db *sqlx.DB) (*Calendar, error) {
	c, err := FindCalendar(ctx, db, s.CalendarId)
	if err != nil {
		return nil, errors.Wrap(err, "FindCalendar")
	}
	return c, nil
}

func (s *Service) Address(ctx context.Context, addr pbaddress.AddressesClient) (*pbaddress.Address, error) {
	return addr.RetrieveAddress(ctx, &pbaddress.RetrieveAddressRequest{AddressId: s.AddressId})
}

func (s *Service) PbPrice() *pbservices.Price {
	return &pbservices.Price{
		Min: s.PriceMin.Float64(),
		Max: s.PriceMax.Float64(),
	}
}

func AvailableServicesInAnyCategory(ctx context.Context, db *sqlx.DB, dow time.Weekday, from, to *pbshared.TimeOfDay, categoryIds ...uint64) ([]*Service, error) {
	cIds, err := calendarIdsAvailablesInRange(ctx, db, dow, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "calendarsWithAvailabilitiesInRange")
	}
	rows, err := db.QueryxContext(ctx,
		`SELECT * FROM services WHERE calendar_id IN($1) AND category_id IN($2)`,
		cIds, categoryIds,
	)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()
	var svcs []*Service
	for rows.Next() {
		svc := &Service{}
		err = rows.StructScan(svc)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		svcs = append(svcs, svc)
	}

	return svcs, nil
}

func (s *Service) ToPb() *pbservices.Service {
	return &pbservices.Service{
		Title:      s.Title,
		UserId:     s.UserId,
		AddressId:  s.AddressId,
		CalendarId: s.CalendarId,

		DurationInMinutes: s.DurationInMinutes,
		Price:             s.PbPrice(),
	}
}
