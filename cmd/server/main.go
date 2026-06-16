package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/AmaanShikalgar/ainyx-users-api/config"
	"github.com/AmaanShikalgar/ainyx-users-api/db"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	database, err := db.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	log.Println("Database connected successfully")

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Ainyx Users API is running",
		})
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"pong": true,
		})
	})

	log.Fatalf("Server error: %v", app.Listen(":"+cfg.AppPort))
}
