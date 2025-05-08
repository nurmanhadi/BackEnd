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

type SubjectService interface {
	AddSubject(request model.SubjectAddRequest) error
	UpdateSubject(subjectId string, request model.SubjectUpdateRequest) error
	FindById(subjectId string) (*entity.Subject, error)
	FindAll() ([]entity.Subject, error)
}
type subjectService struct {
	subjectRepository repository.SubjectRepository
	teacherRepository repository.TeacherRepository
	validation        *validator.Validate
	log               *logrus.Logger
}

func NewSubjectService(
	subjectRepository repository.SubjectRepository,
	teacherRepository repository.TeacherRepository,
	validation *validator.Validate,
	log *logrus.Logger,
) SubjectService {
	return &subjectService{
		subjectRepository: subjectRepository,
		teacherRepository: teacherRepository,
		validation:        validation,
		log:               log,
	}
}
func (s *subjectService) AddSubject(request model.SubjectAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countTeacher, err := s.teacherRepository.CountById(request.TeacherId)
	if err != nil {
		s.log.WithError(err).Error("failed count id to database")
		return err
	}
	if countTeacher < 1 {
		s.log.Warn("teacher not found")
		return exception.NewError(404, "techer not found")
	}
	subject := &entity.Subject{
		TeacherId: request.TeacherId,
		Name:      request.Name,
	}
	if err := s.subjectRepository.Save(*subject); err != nil {
		s.log.WithError(err).Error("failed save subject to database")
	}
	return nil
}
func (s *subjectService) UpdateSubject(subjectId string, request model.SubjectUpdateRequest) error {
	newSubjectId, err := strconv.ParseInt(subjectId, 10, 32)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countSubject, err := s.subjectRepository.CountById(int(newSubjectId))
	if err != nil {
		s.log.WithError(err).Error("failed count subject to database")
		return err
	}
	if countSubject < 1 {
		s.log.Warn("subject not found")
		return exception.NewError(404, "subject not found")
	}
	if request.TeacherId != "" {
		countTeacher, err := s.teacherRepository.CountById(request.TeacherId)
		if err != nil {
			s.log.WithError(err).Error("failed count teacher to database")
			return err
		}
		if countTeacher < 1 {
			s.log.Warn("teacher not found")
			return exception.NewError(404, "techer not found")
		}
	}
	if err := s.subjectRepository.Updates(int(newSubjectId), request); err != nil {
		s.log.WithError(err).Error("failed update subject to database")
	}
	return nil
}
func (s *subjectService) FindById(subjectId string) (*entity.Subject, error) {
	newSubjectId, err := strconv.ParseInt(subjectId, 10, 32)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return nil, err
	}
	subject, err := s.subjectRepository.FindById(int(newSubjectId))
	if err != nil {
		s.log.WithError(err).Warn("subject not found")
		return nil, exception.NewError(404, "subject not found")
	}
	return subject, nil
}
func (s *subjectService) FindAll() ([]entity.Subject, error) {

	subjects, err := s.subjectRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Warn("subject not found")
		return nil, exception.NewError(404, "subject not found")
	}
	return subjects, nil
}
