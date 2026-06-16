package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/AmaanShikalgar/ainyx-users-api/config"
	appdb "github.com/AmaanShikalgar/ainyx-users-api/db"
	dbsqlc "github.com/AmaanShikalgar/ainyx-users-api/db/sqlc"
	"github.com/AmaanShikalgar/ainyx-users-api/internal/handler"
	"github.com/AmaanShikalgar/ainyx-users-api/internal/logger"
	"github.com/AmaanShikalgar/ainyx-users-api/internal/middleware"
	"github.com/AmaanShikalgar/ainyx-users-api/internal/repository"
	"github.com/AmaanShikalgar/ainyx-users-api/internal/routes"
	"github.com/AmaanShikalgar/ainyx-users-api/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	logger.Init(cfg.AppEnv)
	defer logger.Log.Sync()

	logger.Info("starting Ainyx Users API",
		zap.String("env", cfg.AppEnv),
		zap.String("port", cfg.AppPort),
	)

	database, err := appdb.New(cfg)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}
	defer database.Close()

	logger.Info("database connected successfully",
		zap.String("host", cfg.DBHost),
		zap.String("port", cfg.DBPort),
		zap.String("name", cfg.DBName),
	)

	queries := dbsqlc.New(database)

	userRepo := repository.NewUserRepository(queries)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	app := fiber.New()

	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Ainyx Users API is running",
		})
	})

	routes.RegisterRoutes(app, userHandler)

	logger.Info("server starting", zap.String("port", cfg.AppPort))
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		logger.Fatal("server error", zap.Error(err))
	}
}