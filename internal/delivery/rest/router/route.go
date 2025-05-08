package router

import (
	"liva/internal/delivery/rest/handler"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App             *fiber.App
	StudentHandler  handler.StudentHandler
	AuthHandler     handler.AuthHandler
	TeacherHandler  handler.TeacherHandler
	ClassHandler    handler.ClassHandler
	SubjectHandler  handler.SubjectHandler
	ScheduleHandler handler.ScheduleHandler
}

func (r *RouteConfig) Setup() {
	api := r.App.Group("/api")

	student := api.Group("/students")
	student.Get("/", r.StudentHandler.FindAllStudent)
	student.Get("/:studentId", r.StudentHandler.FindStudentById)
	student.Post("/", r.StudentHandler.AddStudent)
	student.Post("/login", r.AuthHandler.StudentLogin)
	student.Put("/:studentId", r.StudentHandler.UpdateStudent)

	teacher := api.Group("teachers")
	teacher.Post("/", r.TeacherHandler.AddTeacher)
	teacher.Put("/:teacherId", r.TeacherHandler.UpdateTeacher)
	teacher.Get("/:teacherId", r.TeacherHandler.FindTeacherById)
	teacher.Get("/", r.TeacherHandler.FindAllTeacher)

	class := api.Group("/classes")
	class.Post("/", r.ClassHandler.AddClass)
	class.Put("/:classId", r.ClassHandler.UpdateClass)
	class.Get("/:classId", r.ClassHandler.FindClassById)
	class.Get("/", r.ClassHandler.FindAllClass)

	subject := api.Group("/subjects")
	subject.Post("/", r.SubjectHandler.AddSubject)
	subject.Get("/", r.SubjectHandler.FindAll)
	subject.Get("/:subjectId", r.SubjectHandler.FindById)
	subject.Put("/:subjectId", r.SubjectHandler.UpdateSubject)

	schadule := api.Group("/schedules")
	schadule.Post("/", r.ScheduleHandler.AddSchedule)
	schadule.Put("/:scheduleId", r.ScheduleHandler.UpdateSchedule)
	schadule.Get("/:scheduleId", r.ScheduleHandler.FindById)
	schadule.Get("/", r.ScheduleHandler.FindAll)
}
