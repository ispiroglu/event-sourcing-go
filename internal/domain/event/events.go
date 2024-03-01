package event

import (
	"write-api/internal/common/eventsourcing"
	"write-api/internal/domain/vo"
)

const (
	EmailChangedEventType    = "EMAIL_CHANGED_EVENT"
	CustomerCreatedEventType = "CUSTOMER_CREATED_EVENT"
	WalletCreatedEventType   = "WALLET_CREATED_EVENT"
)

type CustomerCreatedEvent struct {
	FirstName string
	LastName  string
}

func NewCustomerCreatedEvent(a eventsourcing.AggregateRoot, firstName, lastName string) (*eventsourcing.Event, error) {
	eventData := CustomerCreatedEvent{
		FirstName: firstName,
		LastName:  lastName,
	}

	e := eventsourcing.NewEvent(a, CustomerCreatedEventType)
	if err := e.SetEventData(eventData); err != nil {
		return nil, err
	}

	return e, nil
}

type EmailChangedEvent struct {
	Email string
}

func NewEmailChangedEvent(email string) EmailChangedEvent {
	return EmailChangedEvent{Email: email}
}

type WalletCreatedEvent struct {
	Wallet vo.Wallet
}

func NewWalletCreatedEvent(w vo.Wallet) WalletCreatedEvent {
	return WalletCreatedEvent{Wallet: w}
}
