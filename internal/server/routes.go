package server

import (
	"write-api/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

}

func (s *FiberServer) RegisterCustomerRoutes() {
	s.App.Post("/customer", s.CreateCustomerHandler)

	s.App.Get("/customer/:id", s.GetCustomerHandler)
}

func (s *FiberServer) GetCustomerHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	customer, err := s.GetCustomerById(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{ // Don't mind the status code, it's just for testing
			"error": err.Error(),
		})
	}

	return c.JSON(customer)
}

func (s *FiberServer) CreateCustomerHandler(c *fiber.Ctx) error {
	var r model.CreateCustomerRequest
	c.BodyParser(&r)

	aggregateId, err := s.CreateCustomer(r)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{ // Don't mind the status code, it's just for testing
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":     "Customer created",
		"aggregateId": aggregateId.String(),
	})
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}
