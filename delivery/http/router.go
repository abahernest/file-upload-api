package http

import (
	"github.com/gofiber/fiber/v2"
)

func setupRouter(app *fiber.App, config Config) {
	app.Get("/api/v1/ping", ping)
}
