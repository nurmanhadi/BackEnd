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

type GradeService interface {
	AddGrade(request *model.GradeAddRequest) error
	UpdateGrade(gradeId string, request *model.GradeUpdateRequest) error
	FindById(gradeId string) (*entity.Grade, error)
	FindAll() ([]entity.Grade, error)
}
type gradeService struct {
	gradeRepository   repository.GradeRepository
	studentRepository repository.StudentRepository
	subjectRepository repository.SubjectRepository
	validation        *validator.Validate
	log               *logrus.Logger
}

func NewGradeService(
	gradeRepository repository.GradeRepository,
	studentRepository repository.StudentRepository,
	subjectRepository repository.SubjectRepository,
	validation *validator.Validate,
	log *logrus.Logger,
) GradeService {
	return &gradeService{
		gradeRepository:   gradeRepository,
		studentRepository: studentRepository,
		subjectRepository: subjectRepository,
		validation:        validation,
		log:               log,
	}
}

func (s *gradeService) AddGrade(request *model.GradeAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countStudent, err := s.studentRepository.CountById(request.StudentId)
	if err != nil {
		s.log.WithError(err).Error("failed count student to database")
		return err
	}
	countSubject, err := s.subjectRepository.CountById(request.SubjectId)
	if err != nil {
		s.log.WithError(err).Error("failed count subject to database")
		return err
	}
	if countStudent < 1 {
		s.log.Warn("studetn not found")
		return exception.NewError(404, "student not found")
	} else if countSubject < 1 {
		s.log.Warn("subject not found")
		return exception.NewError(404, "subject not found")
	}
	grade := &entity.Grade{
		StudentId:       request.StudentId,
		SubjectId:       request.SubjectId,
		AttendanceScore: request.AttendanceScore,
		TaskScore:       request.TaskScore,
		MidternScore:    request.MidternScore,
		FinalScore:      request.FinalScore,
	}
	if err := s.gradeRepository.Save(grade); err != nil {
		s.log.WithError(err).Error("failed save grade to database")
		return err
	}
	return nil
}
func (s *gradeService) UpdateGrade(gradeId string, request *model.GradeUpdateRequest) error {
	newGradeId, err := strconv.ParseInt(gradeId, 10, 64)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countGrade, err := s.gradeRepository.CountById(newGradeId)
	if err != nil {
		s.log.WithError(err).Error("failed count grade to database")
		return err
	}
	if countGrade < 1 {
		s.log.Warn("grade not found")
		return exception.NewError(404, "grade not found")
	}
	if request.StudentId != "" {
		countStudent, err := s.studentRepository.CountById(request.StudentId)
		if err != nil {
			s.log.WithError(err).Error("failed count student to database")
			return err
		}
		if countStudent < 1 {
			s.log.Warn("studetn not found")
			return exception.NewError(404, "student not found")
		}
	}
	if request.SubjectId != 0 {
		countSubject, err := s.subjectRepository.CountById(request.SubjectId)
		if err != nil {
			s.log.WithError(err).Error("failed count subject to database")
			return err
		}
		if countSubject < 1 {
			s.log.Warn("subject not found")
			return exception.NewError(404, "subject not found")
		}
	}
	if err := s.gradeRepository.Updates(newGradeId, request); err != nil {
		s.log.WithError(err).Error("failed save grade to database")
		return err
	}
	return nil
}
func (s *gradeService) FindById(gradeId string) (*entity.Grade, error) {
	newGradeId, err := strconv.ParseInt(gradeId, 10, 64)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return nil, err
	}
	grade, err := s.gradeRepository.FindById(newGradeId)
	if err != nil {
		s.log.WithError(err).Warn("grade not found")
		return nil, err
	}
	return grade, nil
}
func (s *gradeService) FindAll() ([]entity.Grade, error) {
	grades, err := s.gradeRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Error("grades not found")
		return nil, err
	}
	return grades, nil
}
