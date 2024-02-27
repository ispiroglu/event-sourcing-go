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

type EmailChangedEvent struct {
	eventsourcing.BaseEvent
	Email string
}

type CustomerCreatedEvent struct {
	eventsourcing.BaseEvent
	FirstName string
	LastName  string
}

type WalletCreatedEvent struct {
	eventsourcing.BaseEvent
	Wallet vo.Wallet
}
