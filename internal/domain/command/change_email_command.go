package command

import "write-api/internal/common/eventsourcing"

type ChangeEmailCommand struct {
	*eventsourcing.BaseCommand
	Email string
}
