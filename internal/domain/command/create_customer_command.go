package command

import "write-api/internal/common/eventsourcing"

type CreateCustomerCommand struct {
	*eventsourcing.BaseCommand
	FirstName string
	LastName  string
}
