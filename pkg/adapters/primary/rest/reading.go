package rest

import "github.com/gofiber/fiber/v2"

func (a *restadapter) handleGetTodos(c *fiber.Ctx) error {
	todos := a.db.Read.Todos()

	return c.Status(fiber.StatusOK).JSON(todos)
}

func (a *restadapter) handleGetDatasets(c *fiber.Ctx) error {
	datasets := a.db.Read.Datasets()

	return c.Status(fiber.StatusOK).JSON(datasets)
}
