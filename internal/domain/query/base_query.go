package query

import (
	uuid "github.com/gofrs/uuid"
)

type BaseQuery struct {
	AggregateId uuid.UUID
}
