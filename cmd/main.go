package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "modernc.org/sqlite"
	"backend-service-api/internal/handlers"
	"backend-service-api/internal/routes"
	"backend-service-api/internal/storage/sqlite"
	"backend-service-api/internal/usecases"
)

func main() {
	app := fiber.New()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	// SQLite connection for users
	dsn := os.Getenv("SQLITE_DSN")
	if dsn == "" {
		dsn = "file:app.db?cache=shared&mode=rwc"
	}
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatalf("sqlite open error: %v", err)
	}
	userRepo, err := sqlite.NewUserSQLiteRepository(db)
	if err != nil {
		log.Fatalf("sqlite repo error: %v", err)
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret"
	}
	authSvc := usecases.NewAuthService(userRepo, jwtSecret, time.Hour*24)
	authHandler := handlers.NewAuthHandler(authSvc)

	routes.Register(app, authHandler, jwtSecret)

	// Optional seed user (set SEED_EMAIL and SEED_PASSWORD)
	if seedEmail := os.Getenv("SEED_EMAIL"); seedEmail != "" {
		if seedPass := os.Getenv("SEED_PASSWORD"); seedPass != "" {
			if _, ok, err := userRepo.GetByEmail(seedEmail); err == nil && !ok {
				if _, err := authSvc.Register(seedEmail, seedPass, os.Getenv("SEED_NAME")); err != nil {
					log.Printf("seed user failed: %v", err)
				} else {
					log.Printf("seed user created: %s", seedEmail)
				}
			}
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

