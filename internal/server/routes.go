package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

}

func (s *FiberServer) RegisterCustomerRoutes() {
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

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}
