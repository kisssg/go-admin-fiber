package routes

import (
	"go-admin/controlllers"
	"go-admin/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controlllers.Register)
	app.Post("/api/login", controlllers.Login)

	app.Use(middlewares.IsAuthenticated)
	app.Get("/api/user", controlllers.User)
	app.Get("/api/logout", controlllers.Logout)

	app.Get("/api/users", controlllers.AllUsers)
	app.Post("/api/users", controlllers.CreateUser)
}
