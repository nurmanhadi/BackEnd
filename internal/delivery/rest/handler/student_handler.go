package handler

import (
	"liva/internal/model"
	"liva/internal/service"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler interface {
	AddStudent(c *fiber.Ctx) error
	FindAllStudent(c *fiber.Ctx) error
	FindStudentById(c *fiber.Ctx) error
	UpdateStudent(c *fiber.Ctx) error
}
type studentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studenService service.StudentService) StudentHandler {
	return &studentHandler{studentService: studenService}
}

func (h *studentHandler) AddStudent(c *fiber.Ctx) error {
	request := new(model.StudentAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.studentService.AddStudent(request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *studentHandler) UpdateStudent(c *fiber.Ctx) error {
	studentId := c.Params("studentId")
	request := new(model.StudentUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.studentService.UpdateStudent(studentId, request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *studentHandler) FindStudentById(c *fiber.Ctx) error {
	studentId := c.Params("studentId")
	student, err := h.studentService.FindStudentById(studentId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, student)
}
func (h *studentHandler) FindAllStudent(c *fiber.Ctx) error {
	students, err := h.studentService.FindAllStudent()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, students)
}
