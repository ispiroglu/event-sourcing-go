package eventsourcing

import (
	"slices"

	uuid "github.com/gofrs/uuid"
)

const (
	appliedEventCap    = 10
	unCommitedEventCap = 10
)

type AggregateRoot interface {
	Apply(e Event) error
	Load(events []Event) error
	RaiseEvent(e Event) error

	GetAggregateId() uuid.UUID
	GetUncommittedEvents() []Event
	GetCommittedEvents() []Event

	ClearUncommittedEvents()
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

func (a *AggregateBase) GetAggregateId() uuid.UUID {
	return a.ID
}

func (a *AggregateBase) GetUncommittedEvents() []Event {
	return a.UncommittedEvents
}

func (a *AggregateBase) GetCommittedEvents() []Event {
	return a.AppliedEvents
}

func (a *AggregateBase) ClearUncommittedEvents() {
	a.UncommittedEvents = a.UncommittedEvents[:0]
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

func (a *AggregateBase) RaiseEvent(e Event) error {
	if e.GetEventId() != a.ID {
		return ErrInvalidAggregateID
	}

	if a.Version >= e.GetVersion() {
		return ErrInvalidEventVersion
	}

	if err := a.eventHandler(e); err != nil {
		return err
	}

	e.SetAggregateType(a.Type) // ? This cannot even get out this function scope? What's the point
	a.Version = e.GetVersion() // ?
	if a.withAppliedEvents {
		a.AppliedEvents = append(a.AppliedEvents, e)
	}

	return nil
}
