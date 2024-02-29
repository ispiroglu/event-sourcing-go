package aggregate

import (
	"write-api/internal/common/eventsourcing"
	"write-api/internal/domain/event"
	"write-api/internal/domain/vo"

	uuid "github.com/gofrs/uuid"
)

const (
	customerAggregateType = "customer"
)

type CustomerAggregate struct {
	*eventsourcing.AggregateBase
	firstName string
	lastName  string
	email     string
	wallet    vo.Wallet
}

// NewCustomer Should this take event as argument?
func NewCustomer(id uuid.UUID) *CustomerAggregate {
	c := &CustomerAggregate{}
	aggregateCfg := &eventsourcing.AggregateConfig{
		Id:                id,
		Type:              customerAggregateType,
		WithAppliedEvents: true,
		Handler:           c.EventHandler,
	}

	return &CustomerAggregate{
		AggregateBase: eventsourcing.NewAggregateBase(aggregateCfg),
		firstName:     "",
		lastName:      "",
		email:         "",
		wallet:        vo.Wallet{},
	}
}

func (c *CustomerAggregate) EventHandler(e eventsourcing.Event) error {

	switch e.(type) {
	case event.CustomerCreatedEvent:
		c.ApplyCustomerCreated(e.(event.CustomerCreatedEvent))
	case event.EmailChangedEvent:
		c.ApplyEmailChanged(e.(event.EmailChangedEvent))
	case event.WalletCreatedEvent:
		c.ApplyWalletCreated(e.(event.WalletCreatedEvent))
	default:
		return eventsourcing.ErrInvalidEventType
	}

	return nil
}

func (c *CustomerAggregate) ApplyCustomerCreated(e event.CustomerCreatedEvent) {
	c.firstName = e.FirstName
	c.lastName = e.LastName
}

func (c *CustomerAggregate) ApplyEmailChanged(event event.EmailChangedEvent) {
	c.email = event.Email
}

func (c *CustomerAggregate) ApplyWalletCreated(event event.WalletCreatedEvent) {
	c.wallet = event.Wallet
}
