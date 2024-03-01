package eventsourcing

import "github.com/gofrs/uuid"

type Command interface {
	GetAggregateId() uuid.UUID
}

type BaseCommand struct {
	AggregateId uuid.UUID
}

func NewBaseCommand(aggregateId uuid.UUID) *BaseCommand {
	return &BaseCommand{AggregateId: aggregateId}
}

func (c *BaseCommand) GetAggregateId() uuid.UUID {
	return c.AggregateId
}
