package handler

import (
	"liva/internal/model"
	"liva/internal/service"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type AttendaceHandler interface {
	AddAttendance(c *fiber.Ctx) error
	UpdateAttendance(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}
type attendaceHandler struct {
	attendaceService service.AttendanceService
}

func NewAttendaceHandler(attendaceService service.AttendanceService) AttendaceHandler {
	return &attendaceHandler{attendaceService: attendaceService}
}
func (h *attendaceHandler) AddAttendance(c *fiber.Ctx) error {
	request := new(model.AttendanceAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.attendaceService.AddAttendance(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *attendaceHandler) UpdateAttendance(c *fiber.Ctx) error {
	attendaceId := c.Params("attendaceId")
	request := new(model.AttendanceUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.attendaceService.UpdateAttendance(attendaceId, *request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *attendaceHandler) FindById(c *fiber.Ctx) error {
	attendaceId := c.Params("attendaceId")
	attendace, err := h.attendaceService.FindById(attendaceId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, attendace)
}
func (h *attendaceHandler) FindAll(c *fiber.Ctx) error {
	attendaces, err := h.attendaceService.FindAll()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, attendaces)
}
