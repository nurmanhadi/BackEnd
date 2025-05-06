package service

import (
	"liva/internal/entity"
	"liva/internal/model"
	"liva/internal/repository"
	"liva/pkg/exception"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type TeacherService interface {
	AddTeacher(request model.TeacherAddRequest) error
	FindAllTeacher() ([]entity.Teacher, error)
	FindTeacherById(teacherId string) (*model.TeacherResponse, error)
	UpdateTeacher(teacherId string, request model.TeacherUpdateRequest) error
}
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
func (s *teacherService) AddTeacher(request model.TeacherAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countNip, err := s.teacherRepository.CountByNip(request.Nip)
	if err != nil {
		s.log.WithError(err).Error("failed count nip to database")
		return err
	}
	if countNip > 0 {
		s.log.Warn("nip alradry exists")
		return exception.NewError(409, "nip already exists")
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed generate hash password")
		return err
	}
	teacher := &entity.Teacher{
		Id:       uuid.NewString(),
		Nip:      request.Nip,
		Name:     strings.ToUpper(request.Name),
		Password: string(newPassword),
	}
	if err := s.teacherRepository.Save(*teacher); err != nil {
		s.log.WithError(err).Error("failed save teacher to database")
		return err
	}
	return nil
}
func (s *teacherService) FindTeacherById(teacherId string) (*model.TeacherResponse, error) {
	teacher, err := s.teacherRepository.FindById(teacherId)
	if err != nil {
		s.log.WithError(err).Warn("teacher not found")
		return nil, exception.NewError(404, "teacher not found")
	}
	response := &model.TeacherResponse{
		Id:   teacher.Id,
		Nip:  teacher.Nip,
		Name: teacher.Name,
	}
	return response, nil
}
func (s *teacherService) FindAllTeacher() ([]entity.Teacher, error) {
	teachers, err := s.teacherRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Warn("teacher not found")
		return nil, err
	}
	return teachers, nil
}
func (s *teacherService) UpdateTeacher(teacherId string, request model.TeacherUpdateRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countId, err := s.teacherRepository.CountById(teacherId)
	if err != nil {
		s.log.WithError(err).Error("failed count id to database")
		return err
	}
	if countId < 1 {
		s.log.Warn("teacher not found")
		return exception.NewError(404, "techer not found")
	}
	if err := s.teacherRepository.Updates(teacherId, &request); err != nil {
		s.log.WithError(err).Error("failed update teacher to database")
		return err
	}
	return nil
}
