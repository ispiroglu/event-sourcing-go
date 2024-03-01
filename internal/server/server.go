package server

import (
	"fmt"
	"log"
	"write-api/internal/application"
	"write-api/internal/domain/command"
	"write-api/internal/domain/query"
	"write-api/internal/infrastructure/repository"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
	*application.CustomerApplicationService
}

func New() *FiberServer {

	// Event storeDB
	eventStoreDb := newEventStoreDb()

	aggregateStore := repository.NewAggregateRepository(eventStoreDb)

	// Query handlers
	getCustomerQueryHandler := query.NewGetCustomerQueryByIdHandler(aggregateStore)

	// Command handlers
	createCustomerCommandHandler := command.NewCreateCustomerCommandHandler(aggregateStore)

	// Application services
	customerApplicationService := application.NewCustomerApplicationService(
		getCustomerQueryHandler,
		createCustomerCommandHandler,
	)

	server := &FiberServer{
		App:                        fiber.New(),
		CustomerApplicationService: customerApplicationService,
	}

	// Register routes
	server.RegisterFiberRoutes()
	server.RegisterCustomerRoutes()

	return server
}

func newEventStoreDb() *esdb.Client {
	cfg, err := esdb.ParseConnectionString("esdb://localhost:2113?tls=false")
	if err != nil {
		log.Fatal(err)
	}

	client, err := esdb.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("client.Config.Address: %v\n", client.Config.Address)

	return client
}
