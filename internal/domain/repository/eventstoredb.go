package repository

type EventStoreDb interface {
	Append(event Event) error
	Load(aggregateId string) ([]Event, error)
}
