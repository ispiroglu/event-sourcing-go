package repository

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"write-api/internal/common/eventsourcing"

	"github.com/EventStore/EventStore-Client-Go/esdb"
)

type EventRepository struct {
	db esdb.Client
}

func NewEventRepository(db esdb.Client) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) SaveEvents(ctx context.Context, streamId string, events []eventsourcing.Event) error {
	eventDataList := make([]esdb.EventData, 0, len(events))
	for _, e := range events {
		eventDataList = append(eventDataList, e.ToEventData())
	}

	_, err := r.db.AppendToStream(
		ctx,
		streamId,
		esdb.AppendToStreamOptions{},
		eventDataList...,
	)
	if err != nil {
		slog.Error("error appending to stream", err)
		return err
	}

	return nil
}

func (r *EventRepository) LoadEvents(ctx context.Context, streamId string) ([]eventsourcing.Event, error) {
	stream, err := r.db.ReadStream(ctx, streamId, esdb.ReadStreamOptions{}, readCount)
	if err != nil {
		slog.Error("error reading stream", err)
		return nil, err
	}
	defer stream.Close()

	events := make([]eventsourcing.Event, 0, 100)
	for {
		re, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			slog.Error("error reading stream", err)
			return nil, err
		}

		events = append(events, *eventsourcing.NewEventFromRecordedEvent(re.Event))
	}

	return events, nil
}
