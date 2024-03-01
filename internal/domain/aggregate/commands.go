package aggregate

import (
	"fmt"
	"write-api/internal/domain/event"
)

// Does aggregate command handlers need commands or just fields?

func (c *CustomerAggregate) CreateCustomer(firstName, lastName string) error {

	// TODO: do the validations here.
	// The checks of the domain. Internal validations.
	if firstName == "" {
		return fmt.Errorf("first name is required")
	}
	if lastName == "" {
		return fmt.Errorf("last name is required")
	}

	event, err := event.NewCustomerCreatedEvent(c, firstName, lastName)
	if err != nil {
		return err
	}

	event.SetMetaData("some metadata")
	return c.Apply(event)
}
