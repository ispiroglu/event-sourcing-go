package eventsourcing

import (
	"github.com/google/uuid"
	"time"
)

// Why did I just declare this xd :)
type Event interface {
	GetEventId() uuid.UUID
	GetAggregateId() uuid.UUID
	GetVersion() int64

	SetAggregateType(string)
	SetVersion(int64)
}

type BaseEvent struct {
	EventID       uuid.UUID
	EventType     string
	Data          []byte
	Timestamp     time.Time
	AggregateType string
	AggregateID   uuid.UUID
	Version       int64
	Metadata      []byte
}

func (e BaseEvent) GetEventId() uuid.UUID {
	return e.EventID
}

func (e BaseEvent) GetAggregateId() uuid.UUID {
	return e.AggregateID
}

func (e BaseEvent) GetVersion() int64 {
	return e.Version
}

func (e BaseEvent) SetAggregateType(_type string) {
	e.AggregateType = _type
}

func (e BaseEvent) SetVersion(version int64) {
	e.Version = version
}
