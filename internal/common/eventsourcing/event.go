package eventsourcing

import (
	"time"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	uuid "github.com/gofrs/uuid"
)

// Why did I just declare this xd :)
type Event interface {
	GetEventId() uuid.UUID
	GetAggregateId() uuid.UUID
	GetVersion() int64

	SetAggregateType(string)
	SetVersion(int64)

	ToEventData() esdb.EventData
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

func NewEventFromRecordedEvent(re *esdb.RecordedEvent) Event {
	return &BaseEvent{
		EventID:     re.EventID,
		EventType:   re.EventType,
		Data:        re.Data,
		Timestamp:   re.CreatedDate,
		AggregateID: uuid.FromStringOrNil(re.StreamID),
		Version:     int64(re.EventNumber),
		Metadata:    re.UserMetadata,
	}
}

func (e *BaseEvent) GetEventId() uuid.UUID {
	return e.EventID
}

func (e *BaseEvent) GetAggregateId() uuid.UUID {
	return e.AggregateID
}

func (e *BaseEvent) GetVersion() int64 {
	return e.Version
}

func (e *BaseEvent) SetAggregateType(_type string) {
	e.AggregateType = _type
}

func (e *BaseEvent) SetVersion(version int64) {
	e.Version = version
}

func (e *BaseEvent) ToEventData() esdb.EventData {
	return esdb.EventData{
		EventID:     e.EventID,
		EventType:   e.EventType,
		ContentType: esdb.JsonContentType,
		Data:        e.Data,
		Metadata:    e.Metadata,
	}
}
