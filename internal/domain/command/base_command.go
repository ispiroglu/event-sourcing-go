package command

import "github.com/google/uuid"

// Some base command
type BaseCommand struct {
	AggregateId uuid.UUID
}
