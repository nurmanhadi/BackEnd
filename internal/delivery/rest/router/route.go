package router

import (
	"liva/internal/delivery/rest/handler"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	StudentHandler handler.StudentHandler
}

func (r *RouteConfig) Setup() {
	api := r.App.Group("/api")

	student := api.Group("/students")
	student.Post("/", r.StudentHandler.AddStudent)
}
