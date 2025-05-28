package handler

import (
	"liva/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler interface {
	CreateAdmin(c *fiber.Ctx) error
	UpdateAdmin(c *fiber.Ctx) error
	GetAdminById(c *fiber.Ctx) error
	GetAllAdmin(c *fiber.Ctx) error
	DeleteAdmin(c *fiber.Ctx) error

	CreateTeacher(c *fiber.Ctx) error
	UpdateTeacher(c *fiber.Ctx) error
	GetTeacherById(c *fiber.Ctx) error
	GetAllTeacher(c *fiber.Ctx) error

	CreateStudent(c *fiber.Ctx) error
	UpdateStudent(c *fiber.Ctx) error
	GetStudentById(c *fiber.Ctx) error
	GetAllStudent(c *fiber.Ctx) error

	CreateClassSubject(c *fiber.Ctx) error
	DeleteClassSubject(c *fiber.Ctx) error

	CreateSchedule(c *fiber.Ctx) error
	UpdateSchedule(c *fiber.Ctx) error
	GetScheduleById(c *fiber.Ctx) error
	GetAllSchedule(c *fiber.Ctx) error

	CreateSubject(c *fiber.Ctx) error
	UpdateSubject(c *fiber.Ctx) error
	GetSubjectById(c *fiber.Ctx) error
	GetAllSubject(c *fiber.Ctx) error

	CreateClass(c *fiber.Ctx) error
	UpdateClass(c *fiber.Ctx) error
	GetClassById(c *fiber.Ctx) error
	GetAllClass(c *fiber.Ctx) error
}
type adminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) AdminHandler {
	return &adminHandler{adminService: adminService}
}
