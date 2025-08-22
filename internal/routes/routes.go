package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend-service-api/internal/handlers"
)

func Register(app *fiber.App, todoHandler *handlers.TodoHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	todos := v1.Group("/todos")

	todos.Get("/", todoHandler.List)
	todos.Post("/", todoHandler.Create)
	todos.Get(":id", todoHandler.Get)
	todos.Put(":id", todoHandler.Update)
	todos.Delete(":id", todoHandler.Delete)
}


