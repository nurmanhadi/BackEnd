package service

import (
	"liva/internal/entity"
	"liva/internal/model"
	"liva/internal/repository"
	"liva/pkg"
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
	userRepository    repository.UserRepository
	validation        *validator.Validate
	log               *logrus.Logger
}

func NewTeacherService(
	teacherRepository repository.TeacherRepository,
	userRepository repository.UserRepository,
	validation *validator.Validate,
	log *logrus.Logger,
) TeacherService {
	return &teacherService{
		teacherRepository: teacherRepository,
		userRepository:    userRepository,
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
		s.log.Warn("nip already exists")
		return exception.NewError(409, "nip already exists")
	}
	teacherId := uuid.NewString()
	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed generate hash password")
		return err
	}
	countIdentifier, err := s.userRepository.CountByIdentifier(request.Nip)
	if err != nil {
		s.log.WithError(err).Error("failed count identifier from database")
		return err
	}
	if countIdentifier > 0 {
		s.log.Warn("identifier already exists")
		return exception.NewError(409, "identifier already exists")
	}
	countReference, err := s.userRepository.CountByReferenceId(teacherId)
	if err != nil {
		s.log.WithError(err).Error("failed count reference_id from database")
		return err
	}
	if countReference > 0 {
		s.log.Warn("reference_id already exists")
		return exception.NewError(409, "reference_id already exists")
	}
	if err := s.userRepository.Save(&entity.User{
		Id:          uuid.NewString(),
		Identifier:  request.Nip,
		Password:    string(newPassword),
		Role:        string(pkg.RoleTeacher),
		ReferenceId: teacherId,
	}); err != nil {
		s.log.WithError(err).Error("failed save user to database")
		return err
	}
	teacher := &entity.Teacher{
		Id:   teacherId,
		Nip:  request.Nip,
		Name: strings.ToUpper(request.Name),
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
