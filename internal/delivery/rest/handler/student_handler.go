package handler

import (
	"liva/internal/model"
	"liva/internal/service"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler interface {
	AddStudent(c *fiber.Ctx) error
}
type studentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studenService service.StudentService) StudentHandler {
	return &studentHandler{studentService: studenService}
}

func (h *studentHandler) AddStudent(c *fiber.Ctx) error {
	request := new(model.StudentRegisterRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.studentService.AddStudent(request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
