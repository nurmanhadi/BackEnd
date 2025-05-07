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

type StudentService interface {
	AddStudent(request *model.StudentAddRequest) error
	FindStudentById(studentId string) (*model.StudentResponse, error)
	FindAllStudent() ([]entity.Student, error)
	UpdateStudent(studentId string, request *model.StudentUpdateRequest) error
}
type studentService struct {
	studentRepository repository.StudentRepository
	validation        *validator.Validate
	log               *logrus.Logger
}

func NewStudentService(studentRepository repository.StudentRepository, validation *validator.Validate, log *logrus.Logger) StudentService {
	return &studentService{
		studentRepository: studentRepository,
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
		return exception.NewError(409, "nis already exists")
	}
	studentId := uuid.NewString()
	newEmail := strings.ToLower(request.Email)
	countEmail, err := s.studentRepository.CountByEmail(newEmail)
	if err != nil {
		s.log.WithError(err).Error("failed count email from database")

		return err
	}
	if countEmail > 0 {
		return exception.NewError(400, "email already exists")
	}
	newName := strings.ToUpper(request.Name)
	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed hash password from bcrypt")

		return err
	}
	student := &entity.Student{
		Id:       studentId,
		Nis:      request.Nis,
		ClassId:  request.ClassId,
		Name:     newName,
		Email:    newEmail,
		Password: string(newPassword),
		Status:   request.Status,
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
		Email:   student.Email,
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
