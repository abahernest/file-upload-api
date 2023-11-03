package http

import (
	"file-upload-api/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Config struct {
	FileRepo domain.FileRepository
}

func RunHttpServer(config Config) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	// setup routes
	setupRouter(app, config)

	return app
}
