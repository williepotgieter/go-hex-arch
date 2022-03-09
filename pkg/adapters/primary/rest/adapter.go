package rest

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/williepotgieter/go-hex-arch/pkg/domain/ports/repository"
)

type restadapter struct {
	app *fiber.App
	db  *repository.DBPort
}

func NewRESTAdapter(db *repository.DBPort) *restadapter {
	app := fiber.New(fiber.Config{
		AppName: "Example API v1",
	})

	app.Use(cors.New(cors.ConfigDefault))
	app.Use(logger.New())

	return &restadapter{app, db}
}

func (a *restadapter) Run(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	p := fmt.Sprintf("127.0.0.1:%v", port)

	a.SetupV1Routes()

	a.app.Listen(p)
}

func (a *restadapter) SetupV1Routes() {
	// API Version 1
	v1 := a.app.Group("/api/v1")

	// Todos
	v1.Get("/todos", a.handleGetTodos)

	// Datasets
	v1.Get("/datasets", a.handleGetDatasets)
}
