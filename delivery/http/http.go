package http

import (
	"github.com/gofiber/fiber/v2"
	"file-upload-api/domain"
)

func ping(c *fiber.Ctx) error {

	return c.JSON(domain.Ping{
		Error: false,
		Msg:   "pong",
	})
}
