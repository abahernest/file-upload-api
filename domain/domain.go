package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Ping struct {
	Error bool   `json:"error" bson:"error"`
	Msg   string `json:"msg" bson:"msg"`
}

type PingRequest struct {
	Message string `json:"message"`
}

type EnvConfig struct {
	DatabaseUrl  string `mapstructure:"DATABASE_URL"`
	DatabaseName string `mapstructure:"DB_NAME"`
}

func HandleError(c *fiber.Ctx, err error) error {
	return c.Status(400).JSON(
		fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
}

func GetSecrets(logger *zap.Logger) {

	_ = godotenv.Load()

}
