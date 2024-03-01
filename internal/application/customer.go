package application

import (
	"context"
	"log/slog"
	"os"
	"write-api/internal/domain/aggregate"
	"write-api/internal/domain/command"
	"write-api/internal/domain/query"
	"write-api/internal/model"

	"github.com/gofrs/uuid"
)

type CustomerApplicationService struct {
	getCustomerQueryHandler *query.GetCustomerQueryByIdHandler

	createCustomerCommandHandler *command.CreateCustomerCommandHandler

	l *slog.Logger
}

func NewCustomerApplicationService(getCustomerQueryHandler *query.GetCustomerQueryByIdHandler, createCustomerCommandHandler *command.CreateCustomerCommandHandler) *CustomerApplicationService {
	return &CustomerApplicationService{
		getCustomerQueryHandler:      getCustomerQueryHandler,
		createCustomerCommandHandler: createCustomerCommandHandler,
		l:                            slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})),
	}
}

func (s *CustomerApplicationService) GetCustomerById(id string) (*aggregate.CustomerAggregate, error) {
	q := query.GetCustomerQueryById{
		BaseQuery: query.BaseQuery{
			AggregateId: uuid.FromStringOrNil(id),
		},
	}

	customer, err := s.getCustomerQueryHandler.Handle(context.Background(), q)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *CustomerApplicationService) CreateCustomer(r model.CreateCustomerRequest) (uuid.UUID, error) {
	aggregateId := uuid.Must(uuid.NewV4())
	cmd := command.NewCreateCustomerCommand(aggregateId, r.FirstName, r.LastName)

	// The context should come from the http request.
	if err := s.createCustomerCommandHandler.Handle(context.Background(), cmd); err != nil {
		s.l.Error("error creating customer", slog.Any("error", err))
		return aggregateId, err
	}

	return aggregateId, nil
}
