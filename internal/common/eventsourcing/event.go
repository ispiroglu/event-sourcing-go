package eventsourcing

import (
	"fmt"
	"time"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	uuid "github.com/gofrs/uuid"
	jsoniter "github.com/json-iterator/go"
)

type Event struct {
	EventID       uuid.UUID
	EventType     string
	Data          []byte
	Timestamp     time.Time
	AggregateType string
	AggregateID   uuid.UUID
	Version       int64
	Metadata      []byte
}

func NewEvent(a AggregateRoot, eventType string) *Event {
	return &Event{
		EventID:       uuid.Must(uuid.NewV4()),
		EventType:     eventType,
		AggregateID:   a.GetAggregateId(),
		AggregateType: a.GetType(),
		Version:       a.GetVersion(),
		Timestamp:     time.Now().UTC(),
	}
}

func NewEventFromRecordedEvent(re *esdb.RecordedEvent) *Event {
	return &Event{
		EventID:     re.EventID,
		EventType:   re.EventType,
		Data:        re.Data,
		Timestamp:   re.CreatedDate,
		AggregateID: uuid.FromStringOrNil(re.StreamID),
		Version:     int64(re.EventNumber),
		Metadata:    re.UserMetadata,
	}
}

func (e *Event) GetEventId() uuid.UUID {
	return e.EventID
}

func (e *Event) GetAggregateId() uuid.UUID {
	return e.AggregateID
}

func (e *Event) GetVersion() int64 {
	return e.Version
}

func (e *Event) SetAggregateType(_type string) {
	e.AggregateType = _type
}

func (e *Event) SetVersion(version int64) {
	e.Version = version
}

func (e *Event) GetEventData(toParse interface{}) error {
	err := jsoniter.Unmarshal(e.Data, toParse)
	fmt.Errorf("%v", err)

	return err
}

func (e *Event) SetEventData(data interface{}) error {
	dataByteArr, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}

	e.Data = dataByteArr
	return nil
}

func (e *Event) GetMetaData(toParse interface{}) error {
	return jsoniter.Unmarshal(e.Metadata, toParse)
}

func (e *Event) SetMetaData(data interface{}) error {
	metaDataByteArr, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}

	e.Metadata = metaDataByteArr
	return nil
}

func (e *Event) ToEventData() esdb.EventData {
	return esdb.EventData{
		EventID:     e.EventID,
		EventType:   e.EventType,
		ContentType: esdb.JsonContentType,
		Data:        e.Data,
		Metadata:    e.Metadata,
	}
}
