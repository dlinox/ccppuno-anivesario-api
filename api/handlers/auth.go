package handlers

import (
	"go-test/database"
	"go-test/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/gofiber/fiber/v2"
)


func verifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func Login(c *fiber.Ctx) error {
	var admin models.Administrator

	var hashedPassword string

	// Decodificar el cuerpo del request
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Comprobar las credenciales
	row := database.DB.QueryRow("SELECT id, name, email, password, role_id FROM administrators WHERE email = ?", admin.Email)
	err := row.Scan(&admin.ID, &admin.Name, &admin.Email, &hashedPassword, &admin.RoleID )
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Credenciales incorrectas",
		})
	}

	if !verifyPassword(hashedPassword, admin.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Contraseña incorrecta",
		})
	}

	// Devolver el administrador como respuesta
	return c.JSON(admin)
}