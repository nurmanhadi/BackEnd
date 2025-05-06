package handler

import (
	"liva/internal/model"
	"liva/internal/service"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	StudentLogin(c *fiber.Ctx) error
}
type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}
func (h *authHandler) StudentLogin(c *fiber.Ctx) error {
	request := new(model.StudentLoginRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	token, err := h.authService.StudentLogin(request)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, token)
}
