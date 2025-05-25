package handler

import (
	"liva/internal/model"
	"liva/internal/service"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler interface {
	AddAdmin(c *fiber.Ctx) error
	UpdateAdmin(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}
type adminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) AdminHandler {
	return &adminHandler{adminService: adminService}
}
func (h *adminHandler) AddAdmin(c *fiber.Ctx) error {
	request := new(model.AdminAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.AddAdmin(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *adminHandler) UpdateAdmin(c *fiber.Ctx) error {
	adminId := c.Params("adminId")
	request := new(model.AdminUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.UpdateAdmin(adminId, *request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *adminHandler) FindById(c *fiber.Ctx) error {
	adminId := c.Params("adminId")
	admin, err := h.adminService.FindById(adminId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, admin)
}
func (h *adminHandler) FindAll(c *fiber.Ctx) error {
	admins, err := h.adminService.FindAll()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, admins)
}
