package repository

import (
	"context"
	"write-api/internal/common/eventsourcing"
)

type IAggregateRepository interface {
	Save(context.Context, eventsourcing.AggregateRoot) error
	Load(context.Context, eventsourcing.AggregateRoot) error
}

type IEventRepository interface {
	SaveEvents(context.Context, string, []eventsourcing.Event) error
	LoadEvents(context.Context, string) ([]eventsourcing.Event, error)
}
