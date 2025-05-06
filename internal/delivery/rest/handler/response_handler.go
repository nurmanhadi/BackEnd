package handler

import "github.com/gofiber/fiber/v2"

func NewResponse[T any](c *fiber.Ctx, code int, data T) error {
	return c.Status(code).JSON(fiber.Map{
		"data": data,
		"path": c.OriginalURL(),
	})
}
