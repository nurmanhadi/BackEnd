package handler

import (
	"liva/internal/model"
	"liva/internal/service"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type GradeHandler interface {
	AddGrade(c *fiber.Ctx) error
	UpdateGrade(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
}
type gradeHandler struct {
	gradeService service.GradeService
}

func NewGradeHandler(gradeService service.GradeService) GradeHandler {
	return &gradeHandler{gradeService: gradeService}
}
func (h *gradeHandler) AddGrade(c *fiber.Ctx) error {
	request := new(model.GradeAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.gradeService.AddGrade(request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *gradeHandler) UpdateGrade(c *fiber.Ctx) error {
	gradeId := c.Params("gradeId")
	request := new(model.GradeUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.gradeService.UpdateGrade(gradeId, request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *gradeHandler) FindAll(c *fiber.Ctx) error {
	grades, err := h.gradeService.FindAll()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, grades)
}
func (h *gradeHandler) FindById(c *fiber.Ctx) error {
	gradeId := c.Params("gradeId")
	grade, err := h.gradeService.FindById(gradeId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, grade)
}
