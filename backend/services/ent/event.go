package ent

import (
	"context"
	"strings"
	"time"

	"github.com/athomecomar/athome/backend/services/ent/schedule"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbshared"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

type Event struct {
	Id         uint64 `json:"id,omitempty"`
	CalendarId uint64 `json:"calendar_id,omitempty"`
	ClaimantId uint64 `json:"claimant_id,omitempty"`

	IsConfirmed bool `json:"is_confirmed,omitempty"`

	OrderId uint64 `json:"order_id,omitempty"`

	DayOfWeek   time.Weekday `json:"day_of_week,omitempty"`
	StartHour   int64        `json:"start_hour,omitempty"`
	EndHour     int64        `json:"end_hour,omitempty"`
	StartMinute int64        `json:"start_minute,omitempty"`
	EndMinute   int64        `json:"end_minute,omitempty"`
}

func eventsWithNonNullIntersectionInRange(ctx context.Context, db *sqlx.DB, dow time.Weekday, from, to *pbshared.TimeOfDay) ([]*Event, error) {
	qr := `
        SELECT * FROM events WHERE
        dow = $1
        AND
        ((start_hour > $2 AND start_minute > $3) AND (start_hour < $4 AND start_minute > $5))
        OR
        ((end_hour > $2 AND end_minute > $3) AND (end_hour < $4 AND end_minute > $5))
    `
	rows, err := db.QueryxContext(ctx, qr, dow, from.GetHour(), from.GetMinute(), to.GetHour(), to.GetMinute())
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()
	var cs []*Event
	for rows.Next() {
		c := &Event{}
		err := rows.StructScan(c)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func EventsToTimeables(es []*Event) (ts []schedule.Scheduleable) {
	for _, e := range es {
		ts = append(ts, e)
	}
	return
}

func EventFromPb(in *pbservices.Event, claimantId uint64, calendarId uint64) (*Event, error) {
	dow, err := DayOfWeekByName(in.GetDow())
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "DayOfWeekByName: %v", err)
	}
	return &Event{
		CalendarId: calendarId,
		ClaimantId: claimantId,

		DayOfWeek:   dow,
		OrderId:     in.GetOrderId(),
		StartHour:   in.GetStart().GetHour(),
		StartMinute: in.GetStart().GetMinute(),
		EndHour:     in.GetEnd().GetHour(),
		EndMinute:   in.GetEnd().GetMinute(),
	}, nil
}

func (e *Event) ToPb() *pbservices.Event {
	return &pbservices.Event{
		Dow:         strings.ToLower(e.DayOfWeek.String()),
		IsConfirmed: e.IsConfirmed,
		ClaimantId:  e.ClaimantId,
		OrderId:     e.OrderId,
		CalendarId:  e.CalendarId,
		Start:       &pbshared.TimeOfDay{Hour: e.StartHour, Minute: e.StartMinute},
		End:         &pbshared.TimeOfDay{Hour: e.EndHour, Minute: e.EndMinute},
	}
}

func (e *Event) GetDayOfWeek() time.Weekday { return e.DayOfWeek }

func (e *Event) GetStartHour() int64   { return e.StartHour }
func (e *Event) GetStartMinute() int64 { return e.StartMinute }

func (e *Event) GetEndHour() int64   { return e.EndHour }
func (e *Event) GetEndMinute() int64 { return e.EndMinute }
