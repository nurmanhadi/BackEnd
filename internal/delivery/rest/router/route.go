package router

import (
	"liva/internal/delivery/rest/handler"
	"liva/internal/delivery/rest/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	StudentHandler    handler.StudentHandler
	AuthHandler       handler.AuthHandler
	TeacherHandler    handler.TeacherHandler
	ClassHandler      handler.ClassHandler
	SubjectHandler    handler.SubjectHandler
	ScheduleHandler   handler.ScheduleHandler
	AttendanceHandler handler.AttendaceHandler
	GradeHandler      handler.GradeHandler
	AdminHandler      handler.AdminHandler
	AuthGuard         *middleware.MiddlewareConfig
}

func (r *RouteConfig) Setup() {
	api := r.App.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", r.AuthHandler.UserLogin)

	admin := api.Group("/admins", r.AuthGuard.Auth())
	admin.Post("/", r.AdminHandler.AddAdmin)
	admin.Put("/:adminId", r.AdminHandler.UpdateAdmin)
	admin.Get("/:adminId", r.AdminHandler.FindById)
	admin.Get("/", r.AdminHandler.FindAll)

	student := api.Group("/students", r.AuthGuard.Auth())
	student.Get("/", r.StudentHandler.FindAllStudent)
	student.Get("/:studentId", r.StudentHandler.FindStudentById)
	student.Post("/", r.StudentHandler.AddStudent)
	student.Put("/:studentId", r.StudentHandler.UpdateStudent)

	teacher := api.Group("teachers", r.AuthGuard.Auth())
	teacher.Post("/", r.TeacherHandler.AddTeacher)
	teacher.Put("/:teacherId", r.TeacherHandler.UpdateTeacher)
	teacher.Get("/:teacherId", r.TeacherHandler.FindTeacherById)
	teacher.Get("/", r.TeacherHandler.FindAllTeacher)

	class := api.Group("/classes", r.AuthGuard.Auth())
	class.Post("/", r.ClassHandler.AddClass)
	class.Put("/:classId", r.ClassHandler.UpdateClass)
	class.Get("/:classId", r.ClassHandler.FindClassById)
	class.Get("/", r.ClassHandler.FindAllClass)

	subject := api.Group("/subjects", r.AuthGuard.Auth())
	subject.Post("/", r.SubjectHandler.AddSubject)
	subject.Get("/", r.SubjectHandler.FindAll)
	subject.Get("/:subjectId", r.SubjectHandler.FindById)
	subject.Put("/:subjectId", r.SubjectHandler.UpdateSubject)

	schadule := api.Group("/schedules", r.AuthGuard.Auth())
	schadule.Post("/", r.ScheduleHandler.AddSchedule)
	schadule.Put("/:scheduleId", r.ScheduleHandler.UpdateSchedule)
	schadule.Get("/:scheduleId", r.ScheduleHandler.FindById)
	schadule.Get("/", r.ScheduleHandler.FindAll)

	attendace := api.Group("/attendances", r.AuthGuard.Auth())
	attendace.Post("/", r.AttendanceHandler.AddAttendance)
	attendace.Put("/:attendaceId", r.AttendanceHandler.UpdateAttendance)
	attendace.Get("/:attendaceId", r.AttendanceHandler.FindById)
	attendace.Get("/", r.AttendanceHandler.FindAll)

	grade := api.Group("/grades", r.AuthGuard.Auth())
	grade.Post("/", r.GradeHandler.AddGrade)
	grade.Get("/", r.GradeHandler.FindAll)
	grade.Get("/:gradeId", r.GradeHandler.FindById)
	grade.Put("/:gradeId", r.GradeHandler.UpdateGrade)
}
