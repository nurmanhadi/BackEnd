package config

import (
	"liva/internal/delivery/rest/handler"
	"liva/internal/delivery/rest/router"
	"liva/internal/repository"
	"liva/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Bootstrap struct {
	App        *fiber.App
	DB         *gorm.DB
	Log        *logrus.Logger
	Validation *validator.Validate
	Viper      *viper.Viper
}

func New(config *Bootstrap) {
	// repository
	studentRepository := repository.NewStudentRepository(config.DB)

	// service
	studentService := service.NewStudentService(studentRepository, config.Validation, config.Log)

	// handler
	studentHandler := handler.NewStudentHandler(studentService)

	// routes
	route := &router.RouteConfig{
		App:            config.App,
		StudentHandler: studentHandler,
	}
	route.Setup()
}
