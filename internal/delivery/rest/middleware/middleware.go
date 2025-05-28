package middleware

import (
	"errors"
	"fmt"
	"liva/internal/service"
	"liva/pkg/exception"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MiddlewareConfig struct {
	App   *fiber.App
	Log   *logrus.Logger
	Viper *viper.Viper
}

func (m *MiddlewareConfig) Application() {
	// monitor
	m.App.Get("/metrics", monitor.New(monitor.Config{Title: "Go User API"}))

	// logger
	// m.App.Use(logger.New(logger.Config{
	// 	Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	// }))

	// compress
	m.App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	m.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Content-Type, Accept, Origin",
	}))
}
func (m *MiddlewareConfig) Auth(role string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString, err := getTokenFromHeader(ctx)
		if err != nil {
			m.Log.WithField("error", err).Warn("failed get token from header")
			return exception.NewError(401, err.Error())
		}
		jwt, err := service.JwtVerifyToken(tokenString, []byte(m.Viper.GetString("jwt.key")))
		if err != nil {
			m.Log.WithField("error", err).Warn("failed verify jwt token")
			return exception.NewError(401, err.Error())
		}
		if jwt.Role != role {
			m.Log.WithField("error", err).Warn("unable role")
			return exception.NewError(401, fmt.Sprintf("only access for %s", role))
		}
		ctx.Locals("id", jwt.Id)
		ctx.Locals("reference_id", jwt.RerefenceId)
		ctx.Locals("role", jwt.Role)
		return ctx.Next()
	}
}
func getTokenFromHeader(c *fiber.Ctx) (string, error) {
	header := c.Get("Authorization")
	if header == "" {
		return "", errors.New("null token Authorization")
	}
	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}
	return parts[1], nil
}
