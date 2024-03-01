package aggregate

import (
	"log/slog"
	"os"
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
	FirstName string
	LastName  string
	Email     string
	Wallet    vo.Wallet

	l *slog.Logger
}

// NewCustomer Should this take event as argument?
func NewCustomer(id uuid.UUID) *CustomerAggregate {
	aggregateCfg := &eventsourcing.AggregateConfig{
		Id:                id,
		Type:              customerAggregateType,
		WithAppliedEvents: true,
	}

	ca := &CustomerAggregate{
		AggregateBase: eventsourcing.NewAggregateBase(aggregateCfg),
		FirstName:     "",
		LastName:      "",
		Email:         "",
		Wallet:        vo.Wallet{},
		l:             slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})),
	}
	ca.SetEventHandler(ca.EventHandler)
	return ca
}

func (c *CustomerAggregate) EventHandler(e *eventsourcing.Event) error {
	switch e.EventType {
	case event.CustomerCreatedEventType:
		err := c.ApplyCustomerCreated(e)
		if err != nil {
			return err
		}
	case event.EmailChangedEventType:
		err := c.ApplyEmailChanged(e)
		if err != nil {
			return err
		}
	case event.WalletCreatedEventType:
		err := c.ApplyWalletCreated(e)
		if err != nil {
			return err
		}
	default:
		return eventsourcing.ErrInvalidEventType
	}

	return nil
}

func (c *CustomerAggregate) ApplyCustomerCreated(evt *eventsourcing.Event) error {
	var e event.CustomerCreatedEvent
	if err := evt.GetEventData(&e); err != nil {
		c.l.Error("Error parsing event", slog.Any("err", err))
		return err
	}

	c.FirstName = e.FirstName
	c.LastName = e.LastName
	return nil
}

func (c *CustomerAggregate) ApplyEmailChanged(evt *eventsourcing.Event) error {
	var e event.EmailChangedEvent
	if err := evt.GetEventData(&e); err != nil {
		c.l.Error("Error parsing event", slog.Any("err", err))
		return err
	}

	c.Email = e.Email
	return nil
}

func (c *CustomerAggregate) ApplyWalletCreated(evt *eventsourcing.Event) error {
	var e event.WalletCreatedEvent
	if err := evt.GetEventData(&e); err != nil {
		c.l.Error("Error parsing event", slog.Any("err", err))
		return err
	}

	c.Wallet = e.Wallet
	return nil
}
