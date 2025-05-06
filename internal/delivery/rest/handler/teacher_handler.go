package handler

import (
	"liva/internal/model"
	"liva/internal/service"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type TeacherHandler interface {
	AddTeacher(c *fiber.Ctx) error
	UpdateTeacher(c *fiber.Ctx) error
	FindTeacherById(c *fiber.Ctx) error
	FindAllTeacher(c *fiber.Ctx) error
}
type teacherHandler struct {
	teacherService service.TeacherService
}

func NewTeacherHandler(teacherService service.TeacherService) TeacherHandler {
	return &teacherHandler{teacherService: teacherService}
}
func (h *teacherHandler) AddTeacher(c *fiber.Ctx) error {
	request := new(model.TeacherAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.teacherService.AddTeacher(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *teacherHandler) UpdateTeacher(c *fiber.Ctx) error {
	teacherId := c.Params("teacherId")
	request := new(model.TeacherUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.teacherService.UpdateTeacher(teacherId, *request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *teacherHandler) FindTeacherById(c *fiber.Ctx) error {
	teacherId := c.Params("teacherId")
	teacher, err := h.teacherService.FindTeacherById(teacherId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, teacher)
}
func (h *teacherHandler) FindAllTeacher(c *fiber.Ctx) error {
	teachers, err := h.teacherService.FindAllTeacher()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, teachers)
}
