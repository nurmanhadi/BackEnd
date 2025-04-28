package config

import (
	"fmt"
	"liva/pkg/exception"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(viper *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      viper.GetString("app.name"),
		Prefork:      viper.GetBool("server.prefork"),
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: errorHandler,
	})
	return app
}

func errorHandler(c *fiber.Ctx, err error) error {
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		var values []string
		for _, fieldErr := range validationErr {
			value := fmt.Sprintf("field %s is %s %s", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
			values = append(values, value)
		}
		str := strings.Join(values, ", ")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": str,
			"path":  c.OriginalURL(),
		})
	}

	if customErr, ok := err.(*exception.CustomError); ok {
		return c.Status(customErr.Code).JSON(fiber.Map{
			"error": customErr.Message,
			"path":  c.OriginalURL(),
		})
	}
	return c.Status(500).JSON(fiber.Map{
		"error": "internal server error",
		"path":  c.OriginalURL(),
	})
}
