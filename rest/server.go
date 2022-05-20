package rest

import (
	"github.com/gofiber/fiber/v2"
)

// Server for REST API
func Server() (app *fiber.App) {
	app = fiber.New()
	return app
}
