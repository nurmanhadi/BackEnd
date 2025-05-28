package router

import (
	"liva/internal/delivery/rest/handler"
	"liva/internal/delivery/rest/middleware"
	"liva/pkg"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	StudentHandler handler.StudentHandler
	AuthHandler    handler.AuthHandler
	TeacherHandler handler.TeacherHandler
	AdminHandler   handler.AdminHandler
	AuthGuard      *middleware.MiddlewareConfig
}

func (r *RouteConfig) Setup() {
	api := r.App.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", r.AuthHandler.UserLogin)

	admin := api.Group("/admins", r.AuthGuard.Auth(string(pkg.RoleAdmin)))
	{
		admin.Post("/", r.AdminHandler.CreateAdmin)
		admin.Put("/:adminId", r.AdminHandler.UpdateAdmin)
		admin.Get("/:adminId", r.AdminHandler.GetAdminById)
		admin.Get("/", r.AdminHandler.GetAllAdmin)
		admin.Delete("/:adminId", r.AdminHandler.DeleteAdmin)

		teacher := admin.Group("/teachers")
		teacher.Post("/", r.AdminHandler.CreateTeacher)
		teacher.Put("/:teacherId", r.AdminHandler.UpdateTeacher)
		teacher.Get("/", r.AdminHandler.GetAllTeacher)
		teacher.Get("/:teacherId", r.AdminHandler.GetTeacherById)

		student := admin.Group("/students")
		student.Post("/", r.AdminHandler.CreateStudent)
		student.Put("/:studentId", r.AdminHandler.UpdateStudent)
		student.Get("/:studentId", r.AdminHandler.GetStudentById)
		student.Get("/", r.AdminHandler.GetAllStudent)

		classSubject := admin.Group("/class-subjects")
		classSubject.Post("/", r.AdminHandler.CreateClassSubject)
		classSubject.Post("/:classSubjectId", r.AdminHandler.DeleteClassSubject)

		schedule := admin.Group("/schedules")
		schedule.Post("/", r.AdminHandler.CreateSchedule)
		schedule.Put("/:studentId", r.AdminHandler.UpdateSchedule)
		schedule.Get("/:studentId", r.AdminHandler.GetScheduleById)
		schedule.Get("/", r.AdminHandler.GetAllSchedule)

		subject := admin.Group("/subjects")
		subject.Post("/", r.AdminHandler.CreateSubject)
		subject.Put("/:studentId", r.AdminHandler.UpdateSubject)
		subject.Get("/:studentId", r.AdminHandler.GetSubjectById)
		subject.Get("/", r.AdminHandler.GetAllSubject)

		class := admin.Group("/classes")
		class.Post("/", r.AdminHandler.CreateClass)
		class.Put("/:studentId", r.AdminHandler.UpdateClass)
		class.Get("/:studentId", r.AdminHandler.GetClassById)
		class.Get("/", r.AdminHandler.GetAllClass)
	}
}
