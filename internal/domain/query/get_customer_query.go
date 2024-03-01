package query

import (
	"context"
	"errors"
	"write-api/internal/domain/aggregate"
	"write-api/internal/infrastructure/repository"
)

type GetCustomerQueryById struct {
	BaseQuery
}

// Do I need an interface?
type GetCustomerQueryByIdHandler struct {
	r *repository.AggregateRepository
}

func NewGetCustomerQueryByIdHandler(r *repository.AggregateRepository) *GetCustomerQueryByIdHandler {
	return &GetCustomerQueryByIdHandler{r: r}
}

func (h GetCustomerQueryByIdHandler) Handle(ctx context.Context, query GetCustomerQueryById) (*aggregate.CustomerAggregate, error) {

	a := aggregate.NewCustomer(query.AggregateId)

	if err := h.r.Load(ctx, a); err != nil {
		return nil, err
	}

	if a.Version == -1 {
		return nil, errors.New("customer not found")
	}

	return a, nil
}
