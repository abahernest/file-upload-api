package domain

import (
	"errors"
	"fmt"
	"net/http"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"github.com/go-playground/validator/v10"
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
	InfuraKey string `mapstructure:"INFURA_API_KEY"`
	InfuraSecret string `mapstructure:"INFURA_API_SECRET"`
}

func HandleError(c *fiber.Ctx, err error, code int, logger *zap.Logger) error {
	logger.Error(err.Error(), zap.Error(err))
	return c.Status(code).JSON(
		fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
}

func GetSecrets(logger *zap.Logger) {

	_ = godotenv.Load()

}


func HandleValidationError(c *fiber.Ctx, err error, logger *zap.Logger) error {

	if _, ok := err.(*validator.InvalidValidationError); ok {

		return HandleError(c, err, 400, logger)
	}

	var errMessage string
	for _, err := range err.(validator.ValidationErrors) {
		errMessage = fmt.Sprintf("enter a valid %v in %v field", err.Kind().String(), err.Field())
		break
	}

	return HandleError(c, errors.New(errMessage), 400, logger)
}

// NewClient creates an http.Client that automatically perform basic auth on each request.
func NewHttpClient(username, password string) *http.Client {
    return &http.Client{
        Transport: &basicAuthTransport{
            Transport:  http.DefaultTransport,
            Username:     username,
            Password: password,
        },
    }
}

// basicAuthTransport decorates each request with a basic auth header.
type basicAuthTransport struct {
    Transport http.RoundTripper
    Username     string
    Password string
}

func (t *basicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    // Add Basic Auth header to the request.
	req.SetBasicAuth(t.Username, t.Password)
	req.Header.Set("authorization", "Basic "+BasicAuth(t.Username, t.Password))

	// send modified request to the underlying Transport
    return t.Transport.RoundTrip(req)
}

func BasicAuth(username, password string) string {
    auth := username + ":" + password
    return base64.StdEncoding.EncodeToString([]byte(auth))
}