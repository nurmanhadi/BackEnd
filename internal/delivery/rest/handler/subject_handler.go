package handler

import (
	"liva/internal/model"
	"liva/internal/service"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type SubjectHandler interface {
	AddSubject(c *fiber.Ctx) error
	UpdateSubject(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}
type subjectHandler struct {
	subjectService service.SubjectService
}

func NewSubjectHandler(subjectService service.SubjectService) SubjectHandler {
	return &subjectHandler{subjectService: subjectService}
}
func (h *subjectHandler) AddSubject(c *fiber.Ctx) error {
	request := new(model.SubjectAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.subjectService.AddSubject(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *subjectHandler) UpdateSubject(c *fiber.Ctx) error {
	subjectId := c.Params("subjectId")
	request := new(model.SubjectUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.subjectService.UpdateSubject(subjectId, *request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *subjectHandler) FindById(c *fiber.Ctx) error {
	subjectId := c.Params("subjectId")
	subject, err := h.subjectService.FindById(subjectId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, subject)
}
func (h *subjectHandler) FindAll(c *fiber.Ctx) error {
	subjects, err := h.subjectService.FindAll()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, subjects)
}
