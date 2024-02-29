package application

import (
	"context"
	"write-api/internal/domain/aggregate"
	"write-api/internal/domain/query"

	"github.com/gofrs/uuid"
)

type CustomerApplicationService struct {
	GetCustomerQueryHandler query.GetCustomerQueryByIdHandler
}

func NewCustomerApplicationService(getCustomerQueryHandler query.GetCustomerQueryByIdHandler) *CustomerApplicationService {
	return &CustomerApplicationService{GetCustomerQueryHandler: getCustomerQueryHandler}
}

func (s CustomerApplicationService) GetCustomerById(id string) (*aggregate.CustomerAggregate, error) {
	query := query.GetCustomerQueryById{
		BaseQuery: query.BaseQuery{
			AggregateId: uuid.FromStringOrNil(id),
		},
	}

	customer, err := s.GetCustomerQueryHandler.Handle(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
