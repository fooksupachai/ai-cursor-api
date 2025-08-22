package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend-service-api/internal/handlers"
	"backend-service-api/internal/middleware"
)

func Register(app *fiber.App, todoHandler *handlers.TodoHandler, authHandler *handlers.AuthHandler, jwtSecret string) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	todos := v1.Group("/todos")

	todos.Get("/", todoHandler.List)
	todos.Post("/", todoHandler.Create)
	todos.Get(":id", todoHandler.Get)
	todos.Put(":id", todoHandler.Update)
	todos.Delete(":id", todoHandler.Delete)

	auth := v1.Group("")
	auth.Post("/login", authHandler.Login)

	protected := v1.Group("")
	protected.Use(middleware.JWTAuth(jwtSecret))
	protected.Get("/me", authHandler.Me)
	protected.Put("/me", authHandler.UpdateProfile)
}


