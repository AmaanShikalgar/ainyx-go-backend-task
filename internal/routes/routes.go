package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/AmaanShikalgar/ainyx-users-api/internal/handler"
)

func RegisterRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	users := app.Group("/users")

	users.Post("/", userHandler.CreateUser)

	users.Get("/", userHandler.GetAllUsers)

	users.Get("/:id", userHandler.GetUserByID)

	users.Put("/:id", userHandler.UpdateUser)

	users.Delete("/:id", userHandler.DeleteUser)
}
