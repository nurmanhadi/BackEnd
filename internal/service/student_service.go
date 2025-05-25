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

type StudentService interface {
	AddStudent(request *model.StudentAddRequest) error
	FindStudentById(studentId string) (*model.StudentResponse, error)
	FindAllStudent() ([]entity.Student, error)
	UpdateStudent(studentId string, request *model.StudentUpdateRequest) error
}
type studentService struct {
	studentRepository repository.StudentRepository
	userRepository    repository.UserRepository
	validation        *validator.Validate
	log               *logrus.Logger
}

func NewStudentService(studentRepository repository.StudentRepository, userRepository repository.UserRepository, validation *validator.Validate, log *logrus.Logger) StudentService {
	return &studentService{
		studentRepository: studentRepository,
		userRepository:    userRepository,
		log:               log,
		validation:        validation,
	}
}
func (s *studentService) AddStudent(request *model.StudentAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countNis, err := s.studentRepository.CountByNis(request.Nis)
	if err != nil {
		s.log.WithError(err).Error("failed count nis from database")
		return err
	}
	if countNis > 0 {
		s.log.Warn("nis already exists")
		return exception.NewError(409, "nis already exists")
	}
	studentId := uuid.NewString()
	newName := strings.ToUpper(request.Name)
	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed hash password from bcrypt")
		return err
	}
	countIdentifier, err := s.userRepository.CountByIdentifier(request.Nis)
	if err != nil {
		s.log.WithError(err).Error("failed count identifier from database")
		return err
	}
	if countIdentifier > 0 {
		s.log.Warn("identifier already exists")
		return exception.NewError(409, "identifier already exists")
	}
	countReference, err := s.userRepository.CountByReferenceId(studentId)
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
		Identifier:  request.Nis,
		Password:    string(newPassword),
		Role:        string(pkg.RoleStudent),
		ReferenceId: studentId,
	}); err != nil {
		s.log.WithError(err).Error("failed save user to database")
		return err
	}
	student := &entity.Student{
		Id:      studentId,
		Nis:     request.Nis,
		ClassId: request.ClassId,
		Name:    newName,
		Status:  request.Status,
	}
	err = s.studentRepository.Save(*student)
	if err != nil {
		s.log.WithError(err).Error("failed save student to database")
		return err
	}
	return nil
}
func (s *studentService) FindStudentById(studentId string) (*model.StudentResponse, error) {
	student, err := s.studentRepository.FindById(studentId)
	if err != nil {
		s.log.WithError(err).Warn("student not found")
		return nil, exception.NewError(404, "student not found")
	}
	response := &model.StudentResponse{
		Id:      student.Id,
		Nis:     student.Nis,
		ClassId: student.ClassId,
		Name:    student.Name,
		Status:  student.Status,
	}
	return response, nil
}
func (s *studentService) FindAllStudent() ([]entity.Student, error) {
	students, err := s.studentRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Warn("student not found")
		return nil, exception.NewError(404, "student not found")
	}
	return students, nil
}
func (s *studentService) UpdateStudent(studentId string, request *model.StudentUpdateRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countId, err := s.studentRepository.CountById(studentId)
	if err != nil {
		s.log.WithError(err).Error("failed count id from database")
		return err
	}
	if countId < 1 {
		return exception.NewError(404, "student not found")
	}
	err = s.studentRepository.Updates(studentId, request)
	if err != nil {
		s.log.WithError(err).Error("failed update student to database")
		return err
	}
	return nil
}
