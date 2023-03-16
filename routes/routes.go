package routes

import (
	"go-admin/controlllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controlllers.Register)
	app.Post("/api/login", controlllers.Login)
	app.Get("/api/user", controlllers.User)
	app.Get("/api/logout", controlllers.Logout)
}
