package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"backend-service-api/internal/handlers"
	"backend-service-api/internal/routes"
	"backend-service-api/internal/storage/memory"
	"backend-service-api/internal/usecases"
)

func main() {
	app := fiber.New()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	// Wire dependencies
	repo := memory.NewTodoMemoryRepository()
	service := usecases.NewTodoService(repo)
	h := handlers.NewTodoHandler(service)
	routes.Register(app, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

