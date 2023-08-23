package main

import (
	"github.com/gofiber/fiber/v2"
	"go-test/api/routes"
	"go-test/database"
)

func main() {
	app := fiber.New()
	// Iniciar la base de datos
	database.InitDB()
	defer database.DB.Close()
	// Configurar rutas
	routes.AuthRoutes(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
}
