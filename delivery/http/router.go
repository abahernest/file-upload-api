package http

import (
	fileRoutes "file-upload-api/delivery/http/files"
	fileUsecase "file-upload-api/application/files"

	"github.com/gofiber/fiber/v2"
)

func setupRouter(app *fiber.App, config Config) {
	app.Get("/api/v1/ping", ping)

	v1RouteGroup := app.Group("/api/v1/")

	fileRouteGroup := v1RouteGroup.Group("/file")

	fileRoutes.New(fileRouteGroup, v1RouteGroup, fileUsecase.New(config.FileRepo))
}
