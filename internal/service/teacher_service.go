package service

import (
	"liva/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type TeacherService interface{}
type teacherService struct {
	teacherRepository repository.TeacherRepository
	validation        *validator.Validate
	log               *logrus.Logger
}

func NewTeacherService(
	teacherRepository repository.TeacherRepository,
	validation *validator.Validate,
	log *logrus.Logger,
) TeacherService {
	return &teacherService{
		teacherRepository: teacherRepository,
		validation:        validation,
		log:               log,
	}
}
