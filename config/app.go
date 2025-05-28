package config

import (
	"liva/internal/delivery/rest/handler"
	"liva/internal/delivery/rest/middleware"
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

func New(config *Bootstrap) { // dependency injection
	// repository
	userRepository := repository.NewUserRepository(config.DB)
	adminRepository := repository.NewAdminRepository(config.DB)
	studentRepository := repository.NewStudentRepository(config.DB)
	teacherRepository := repository.NewTeacherRepository(config.DB)
	classRepository := repository.NewClassRepository(config.DB)
	subjectRepository := repository.NewSubjectRepository(config.DB)
	classSubjectRepository := repository.NewClassSubjectRepository(config.DB)
	scheduleRepository := repository.NewScheduleRepository(config.DB)
	// attendaceRepository := repository.NewAttendaceRepository(config.DB)
	// gradeRepository := repository.NewGradeRepository(config.DB)

	// service
	adminService := service.NewAdminService(
		adminRepository, userRepository, teacherRepository,
		studentRepository, classRepository, subjectRepository,
		classSubjectRepository, scheduleRepository, config.Validation, config.Log,
	)
	studentService := service.NewStudentService(studentRepository, userRepository, config.Validation, config.Log)
	authService := service.NewAuthService(userRepository, config.Validation, config.Log, config.Viper)
	// teacherService := service.NewTeacherService(teacherRepository, userRepository, config.Validation, config.Log)
	// handler
	adminHandler := handler.NewAdminHandler(adminService)
	studentHandler := handler.NewStudentHandler(studentService)
	authHandler := handler.NewAuthHandler(authService)
	// teacherHadnler := handler.NewTeacherHandler(teacherService)

	// middleware
	middleware := middleware.MiddlewareConfig{
		App:   config.App,
		Log:   config.Log,
		Viper: config.Viper,
	}
	middleware.Application()

	// routes
	route := &router.RouteConfig{
		App:            config.App,
		StudentHandler: studentHandler,
		AuthHandler:    authHandler,
		// TeacherHandler: teacherHadnler,
		AdminHandler: adminHandler,
		AuthGuard:    &middleware,
	}
	route.Setup()
}
