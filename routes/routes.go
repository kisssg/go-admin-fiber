package routes

import (
	"go-admin/controlllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controlllers.Register)
	app.Post("/api/login", controlllers.Login)
}
