package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/podanypepa/cursor-pagination/rest/controller"
)

// Server for REST API
func Server() (app *fiber.App) {
	app = fiber.New()
	app.Get("/user", controller.User)
	return app
}
