package main

import (
"log"
"os"

"github.com/gofiber/fiber/v2"
)

func main() {
app := fiber.New()

app.Get("/healthz", func(c *fiber.Ctx) error {
return c.SendString("Hello World")
})

port := os.Getenv("PORT")
if port == "" {
port = "8080"
}

log.Printf("Starting server on :%s", port)
if err := app.Listen(":" + port); err != nil {
log.Fatalf("failed to start server: %v", err)
}
}

