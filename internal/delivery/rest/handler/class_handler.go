package handler

import (
	"liva/internal/model"
	"liva/internal/service"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type ClassHandler interface {
	AddClass(c *fiber.Ctx) error
	UpdateClass(c *fiber.Ctx) error
	FindClassById(c *fiber.Ctx) error
	FindAllClass(c *fiber.Ctx) error
}
type classHandler struct {
	classService service.ClassService
}

func NewClassHandler(classService service.ClassService) ClassHandler {
	return &classHandler{classService: classService}
}
func (h *classHandler) AddClass(c *fiber.Ctx) error {
	request := new(model.ClassAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.classService.AddClass(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *classHandler) UpdateClass(c *fiber.Ctx) error {
	classId := c.Params("classId")
	request := new(model.ClassUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.classService.UpdateClass(classId, *request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *classHandler) FindClassById(c *fiber.Ctx) error {
	classId := c.Params("classId")
	class, err := h.classService.FindClassById(classId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, class)
}
func (h *classHandler) FindAllClass(c *fiber.Ctx) error {
	classes, err := h.classService.FindAllClass()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, classes)
}
