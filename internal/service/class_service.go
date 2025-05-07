package service

import (
	"liva/internal/entity"
	"liva/internal/model"
	"liva/internal/repository"
	"liva/pkg/exception"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ClassService interface {
	AddClass(request model.ClassAddRequest) error
	UpdateClass(classId string, request model.ClassUpdateRequest) error
	FindClassById(classId string) (*entity.Class, error)
	FindAllClass() ([]entity.Class, error)
}
type classService struct {
	classRepository   repository.ClassRepository
	teacherRepository repository.TeacherRepository
	validation        *validator.Validate
	log               *logrus.Logger
}

func NewClassService(classRepository repository.ClassRepository, teacherRepository repository.TeacherRepository, validation *validator.Validate, log *logrus.Logger) ClassService {
	return &classService{
		classRepository:   classRepository,
		teacherRepository: teacherRepository,
		validation:        validation,
		log:               log,
	}
}
func (s *classService) AddClass(request model.ClassAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countTeacher, err := s.teacherRepository.CountById(request.TeacherId)
	if err != nil {
		s.log.WithError(err).Error("failed count id to database")
		return err
	}
	countTeacherId, err := s.classRepository.CountByTeacherId(request.TeacherId)
	if err != nil {
		s.log.WithError(err).Error("failed count id to database")
		return err
	}
	if countTeacher < 1 {
		s.log.Warn("teacher not found")
		return exception.NewError(404, "teacher not found")
	} else if countTeacherId > 0 {
		s.log.Warn("teacher already exists")
		return exception.NewError(409, "teacher already exists")
	}
	class := &entity.Class{
		TeacherId: request.TeacherId,
		Name:      request.Name,
	}
	if err := s.classRepository.Save(*class); err != nil {
		s.log.WithError(err).Error("failed save to database")
		return err
	}
	return nil
}
func (s *classService) UpdateClass(classId string, request model.ClassUpdateRequest) error {
	newClassId, err := strconv.ParseInt(classId, 10, 32)
	if err != nil {
		s.log.WithError(err).Error("failed parse to int")
		return err
	}
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countClass, err := s.classRepository.CountById(int(newClassId))
	if err != nil {
		s.log.WithError(err).Error("failed count id to database")
		return err
	}
	if request.TeacherId != "" {
		countTeacher, err := s.teacherRepository.CountById(request.TeacherId)
		if err != nil {
			s.log.WithError(err).Error("failed count id to database")
			return err
		}
		if countTeacher < 1 {
			s.log.Warn("teacher not found")
			return exception.NewError(404, "teacher not found")
		}
	}
	if countClass < 1 {
		s.log.Warn("class not found")
		return exception.NewError(404, "class not found")
	}
	if err := s.classRepository.Updates(int(newClassId), request); err != nil {
		s.log.WithError(err).Error("failed update to database")
		return err
	}
	return nil
}
func (s *classService) FindClassById(classId string) (*entity.Class, error) {
	newClassId, err := strconv.ParseInt(classId, 10, 32)
	if err != nil {
		s.log.WithError(err).Error("failed parse to int")
		return nil, err
	}
	class, err := s.classRepository.FindById(int(newClassId))
	if err != nil {
		s.log.WithError(err).Warn("class not found")
		return nil, exception.NewError(404, "class not found")
	}
	return class, nil
}
func (s *classService) FindAllClass() ([]entity.Class, error) {
	classes, err := s.classRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Warn("failed find all class to database")
		return nil, err
	}
	return classes, nil
}
