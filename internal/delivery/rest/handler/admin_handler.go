package handler

import (
	"liva/internal/model"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

// ==========================
// admin section
// ==========================
func (h *adminHandler) CreateAdmin(c *fiber.Ctx) error {
	request := new(model.AdminAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.CreateAdmin(*request); err != nil {
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
func (h *adminHandler) GetAdminById(c *fiber.Ctx) error {
	adminId := c.Params("adminId")
	admin, err := h.adminService.GetAdminById(adminId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, admin)
}
func (h *adminHandler) GetAllAdmin(c *fiber.Ctx) error {
	admins, err := h.adminService.GetAllAdmin()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, admins)
}
func (h *adminHandler) DeleteAdmin(c *fiber.Ctx) error {
	adminId := c.Params("adminId")
	err := h.adminService.DeleteAdmin(adminId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}

// ==========================
// teacher section
// ==========================
func (h *adminHandler) CreateTeacher(c *fiber.Ctx) error {
	request := new(model.TeacherAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.CreateTeacher(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *adminHandler) UpdateTeacher(c *fiber.Ctx) error {
	teacherId := c.Params("teacherId")
	request := new(model.TeacherUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.UpdateTeacher(teacherId, *request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *adminHandler) GetTeacherById(c *fiber.Ctx) error {
	teacherId := c.Params("teacherId")
	teacher, err := h.adminService.GetTeacherById(teacherId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, teacher)
}
func (h *adminHandler) GetAllTeacher(c *fiber.Ctx) error {
	teachers, err := h.adminService.GetAllTeacher()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, teachers)
}

// ==========================
// student section
// ==========================
func (h *adminHandler) CreateStudent(c *fiber.Ctx) error {
	request := new(model.StudentAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.CreateStudent(request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *adminHandler) UpdateStudent(c *fiber.Ctx) error {
	studentId := c.Params("studentId")
	request := new(model.StudentUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.UpdateStudent(studentId, request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *adminHandler) GetStudentById(c *fiber.Ctx) error {
	studentId := c.Params("studentId")
	student, err := h.adminService.GetStudentById(studentId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, student)
}
func (h *adminHandler) GetAllStudent(c *fiber.Ctx) error {
	students, err := h.adminService.GetAllStudent()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, students)
}

// ==========================
// class-subject section
// ==========================
func (h *adminHandler) CreateClassSubject(c *fiber.Ctx) error {
	request := new(model.ClassSubjectAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.CreateClassSubject(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *adminHandler) DeleteClassSubject(c *fiber.Ctx) error {
	classSubjectId := c.Params("classSubjectId")
	err := h.adminService.DeleteClassSubject(classSubjectId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}

// ==========================
// schedule section
// ==========================
func (h *adminHandler) CreateSchedule(c *fiber.Ctx) error {
	request := new(model.ScheduleAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.CreateSchedule(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *adminHandler) UpdateSchedule(c *fiber.Ctx) error {
	scheduleId := c.Params("scheduleId")
	request := new(model.ScheduleUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.UpdateSchedule(scheduleId, *request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *adminHandler) GetScheduleById(c *fiber.Ctx) error {
	scheduleId := c.Params("scheduleId")
	schedule, err := h.adminService.GetScheduleById(scheduleId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, schedule)
}
func (h *adminHandler) GetAllSchedule(c *fiber.Ctx) error {
	schedules, err := h.adminService.GetAllSchedule()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, schedules)
}

// ==========================
// subject section
// ==========================
func (h *adminHandler) CreateSubject(c *fiber.Ctx) error {
	request := new(model.SubjectAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.CreateSubject(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *adminHandler) UpdateSubject(c *fiber.Ctx) error {
	subjectId := c.Params("subjectId")
	request := new(model.SubjectUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.UpdateSubject(subjectId, *request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *adminHandler) GetSubjectById(c *fiber.Ctx) error {
	subjectId := c.Params("subjectId")
	subject, err := h.adminService.GetSubjectById(subjectId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, subject)
}
func (h *adminHandler) GetAllSubject(c *fiber.Ctx) error {
	subjects, err := h.adminService.GetAllSubject()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, subjects)
}

// ==========================
// class section
// ==========================
func (h *adminHandler) CreateClass(c *fiber.Ctx) error {
	request := new(model.ClassAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.CreateClass(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *adminHandler) UpdateClass(c *fiber.Ctx) error {
	classId := c.Params("classId")
	request := new(model.ClassUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.adminService.UpdateClass(classId, *request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *adminHandler) GetClassById(c *fiber.Ctx) error {
	classId := c.Params("classId")
	class, err := h.adminService.GetClassById(classId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, class)
}
func (h *adminHandler) GetAllClass(c *fiber.Ctx) error {
	classes, err := h.adminService.GetAllClass()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, classes)
}
