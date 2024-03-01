package command

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"write-api/internal/common/eventsourcing"
	"write-api/internal/domain/aggregate"
	"write-api/internal/infrastructure/repository"

	"github.com/gofrs/uuid"
)

type CreateCustomerCommand struct {
	*eventsourcing.BaseCommand
	FirstName string
	LastName  string
}

func NewCreateCustomerCommand(aggregateId uuid.UUID, firstName, lastName string) *CreateCustomerCommand {
	return &CreateCustomerCommand{
		BaseCommand: eventsourcing.NewBaseCommand(aggregateId),
		FirstName:   firstName,
		LastName:    lastName,
	}
}

type CreateCustomerCommandHandler struct {
	r *repository.AggregateRepository
	l *slog.Logger
}

func NewCreateCustomerCommandHandler(r *repository.AggregateRepository) *CreateCustomerCommandHandler {
	return &CreateCustomerCommandHandler{
		r: r,
		l: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})),
	}
}

func (h *CreateCustomerCommandHandler) Handle(ctx context.Context, c *CreateCustomerCommand) error {
	a := aggregate.NewCustomer(c.AggregateId)

	// TODO: There needs to be an existing control
	// The checks of our world, not for domain.

	if err := a.CreateCustomer(c.FirstName, c.LastName); err != nil {
		h.l.Error("error creating customer", slog.Any("error", err))
		return err
	}

	fmt.Println("asda")
	return h.r.Save(ctx, a)
}
