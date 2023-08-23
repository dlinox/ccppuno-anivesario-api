package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-test/api/handlers"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/login", handlers.Login)
}