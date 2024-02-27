package eventsourcing

import (
	"github.com/google/uuid"
	"slices"
)

const (
	appliedEventCap    = 10
	unCommitedEventCap = 10
)

type AggregateRoot interface {
	Apply(e BaseEvent) error
	Load(e BaseEvent) error
	// When(e BaseEvent) error
}

type EventHandler func(e Event) error

type AggregateBase struct {
	ID                uuid.UUID
	Version           int64
	AppliedEvents     []Event
	UncommittedEvents []Event
	Type              string
	withAppliedEvents bool
	eventHandler      EventHandler
}

type AggregateConfig struct {
	Id                uuid.UUID
	Type              string
	WithAppliedEvents bool
	Handler           EventHandler
}

func NewAggregateBase(cfg *AggregateConfig) *AggregateBase {
	if cfg.Handler == nil {
		return nil
	}

	return &AggregateBase{
		ID:                cfg.Id,
		Version:           -1,
		AppliedEvents:     make([]Event, 0, appliedEventCap),
		UncommittedEvents: make([]Event, 0, unCommitedEventCap),
		Type:              cfg.Type,
		withAppliedEvents: cfg.WithAppliedEvents,
		eventHandler:      cfg.Handler,
	}
}

func (a *AggregateBase) Load(events []Event) error {
	for _, event := range events {
		if event.GetAggregateId() != a.ID {
			return ErrInvalidAggregateID
		}

		if err := a.eventHandler(event); err != nil {
			return err
		}

		a.Version++
	}
	slices.Concat(a.AppliedEvents, events)
	return nil
}

func (a *AggregateBase) Apply(e Event) error {
	if e.GetEventId() != a.ID {
		return ErrInvalidAggregateID
	}

	if err := a.eventHandler(e); err != nil {
		return err
	}

	a.Version++
	a.UncommittedEvents = append(a.UncommittedEvents, e)
	e.SetAggregateType(a.Type)
	e.SetVersion(a.Version)

	return nil
}
