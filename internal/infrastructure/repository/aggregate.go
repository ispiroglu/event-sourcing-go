package repository

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"math"
	"write-api/internal/common/eventsourcing"

	"github.com/EventStore/EventStore-Client-Go/esdb"
)

const readCount = math.MaxUint64

type AggregateRepository struct {
	db *esdb.Client
}

func NewAggregateRepository(db *esdb.Client) *AggregateRepository {
	return &AggregateRepository{db: db}
}

// Should this take aggregateRoot or aggregateBase?
func (r *AggregateRepository) Load(ctx context.Context, a eventsourcing.AggregateRoot) error {
	stream, err := r.db.ReadStream(ctx, a.GetAggregateId().String(), esdb.ReadStreamOptions{}, readCount)
	if err != nil {
		slog.Error("error reading stream", err)
		return err
	}
	defer stream.Close()

	for {
		re, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if errors.Is(err, esdb.ErrStreamNotFound) {
			slog.Error("stream not found", err)
			return err
		}

		e := eventsourcing.NewEventFromRecordedEvent(re.Event)

		if err := a.RaiseEvent(e); err != nil {
			slog.Error("error raising event", err)
			return err
		}
	}

	return nil
}

func (r *AggregateRepository) Save(ctx context.Context, a eventsourcing.AggregateRoot) error {
	if len(a.GetUncommittedEvents()) == 0 {
		return nil
	}

	eventDataList := make([]esdb.EventData, 0, len(a.GetUncommittedEvents()))
	for _, e := range a.GetUncommittedEvents() {
		eventDataList = append(eventDataList, e.ToEventData())
	}

	if len(a.GetCommittedEvents()) == 0 {
		_, err := r.db.AppendToStream(
			ctx,
			a.GetAggregateId().String(),
			esdb.AppendToStreamOptions{ExpectedRevision: esdb.NoStream{}},
			eventDataList...,
		)
		if err != nil {
			slog.Error("error appending to stream", err)
			return err
		}

		return nil
	}

	readStream, err := r.db.ReadStream(
		ctx,
		a.GetAggregateId().String(),
		esdb.ReadStreamOptions{Direction: esdb.Backwards, From: esdb.End{}},
		readCount,
	)
	if err != nil {
		slog.Error("error reading stream", err)
		return err
	}
	defer readStream.Close()

	lastEvent, err := readStream.Recv()
	if err != nil {
		slog.Error("error reading stream", err)
		return err
	}

	expectedRevision := esdb.Revision(lastEvent.OriginalEvent().EventNumber)
	_, err = r.db.AppendToStream(
		ctx,
		a.GetAggregateId().String(),
		esdb.AppendToStreamOptions{ExpectedRevision: expectedRevision},
		eventDataList...,
	)
	if err != nil {
		slog.Error("error appending to stream", err)
		return err
	}

	a.ClearUncommittedEvents()
	return nil
}
