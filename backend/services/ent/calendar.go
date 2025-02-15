package ent

import (
	"context"
	"time"

	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbshared"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Calendar struct {
	Id      uint64 `json:"id,omitempty"`
	UserId  uint64 `json:"user_id,omitempty"`
	GroupId uint64 `json:"group_id,omitempty"`

	Name string `json:"name,omitempty"`
}

func CalendarsByUserId(ctx context.Context, db *sqlx.DB, uid uint64) ([]*Calendar, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM calendars WHERE user_id=$1`, uid)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()
	var cs []*Calendar
	for rows.Next() {
		c := &Calendar{}
		err := rows.StructScan(c)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func CalendarFromPb(c *pbservices.Calendar, uid uint64) *Calendar {
	return &Calendar{
		Name:    c.Name,
		UserId:  uid,
		GroupId: c.GroupId,
	}
}

func (c *Calendar) ToPb() *pbservices.Calendar {
	return &pbservices.Calendar{
		Name:    c.Name,
		GroupId: c.GroupId,
		UserId:  c.UserId,
	}
}

func FindCalendar(ctx context.Context, db *sqlx.DB, id uint64) (*Calendar, error) {
	c := &Calendar{}
	row := storeql.Where(ctx, db, c, `id=$1`, id)
	err := row.StructScan(c)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return c, nil
}

func (c *Calendar) Availabilities(ctx context.Context, db *sqlx.DB) ([]*Availability, error) {
	rows, err := storeql.WhereMany(ctx, db, &Availability{}, `calendar_id=$1`, c.Id)
	if err != nil {
		return nil, errors.Wrap(err, "WhereMany")
	}
	defer rows.Close()
	var avs []*Availability
	for rows.Next() {
		av := &Availability{}
		err := rows.StructScan(av)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		avs = append(avs, av)
	}
	return avs, nil
}

func (c *Calendar) Events(ctx context.Context, db *sqlx.DB) ([]*Event, error) {
	rows, err := storeql.WhereMany(ctx, db, &Event{}, `calendar_id=$1`, c.Id)
	if err != nil {
		return nil, errors.Wrap(err, "WhereMany")
	}
	defer rows.Close()
	var avs []*Event
	for rows.Next() {
		av := &Event{}
		err := rows.StructScan(av)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		avs = append(avs, av)
	}
	return avs, nil
}

func calendarIdsAvailablesInRange(ctx context.Context, db *sqlx.DB, dow time.Weekday, from, to *pbshared.TimeOfDay) ([]uint64, error) {
	avs, err := availabilitiesContainingRange(ctx, db, dow, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "availabilitiesContainingRange")
	}
	evs, err := eventsWithNonNullIntersectionInRange(ctx, db, dow, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "eventsWithNonNullIntersectionInRange")
	}
	var cIds []uint64
	for _, av := range avs {
		var isAvailable = true
		for _, ev := range evs {
			if ev.CalendarId == av.CalendarId {
				isAvailable = false
			}
		}
		if isAvailable {
			cIds = append(cIds, av.CalendarId)
		}
	}
	return cIds, nil
}

func CalendarsAvailablesInRange(ctx context.Context, db *sqlx.DB, dow time.Weekday, from, to *pbshared.TimeOfDay) ([]*Calendar, error) {
	ids, err := calendarIdsAvailablesInRange(ctx, db, dow, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "calendarIdsAvailablesInRange")
	}
	rows, err := storeql.WhereMany(ctx, db, &Calendar{}, `id IN ($1)`, ids)
	if err != nil {
		return nil, errors.Wrap(err, "WhereMany")
	}
	defer rows.Close()
	var cs []*Calendar
	for rows.Next() {
		c := &Calendar{}
		err := rows.StructScan(c)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func AvailabilitiesByUserGroup(ctx context.Context, db *sqlx.DB, uid, gid uint64) ([]*Availability, error) {
	rows, err := storeql.WhereMany(ctx, db, &Availability{}, `calendar_id IN (SELECT id FROM calendars WHERE user_id=$1 AND group_id=$2)`, uid, gid)
	if err != nil {
		return nil, errors.Wrap(err, "WhereMany")
	}
	defer rows.Close()
	var avs []*Availability
	for rows.Next() {
		av := &Availability{}
		err := rows.StructScan(av)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		avs = append(avs, av)
	}
	return avs, nil
}
