package service

import (
	"liva/internal/entity"
	"liva/internal/model"
	"liva/internal/repository"
	"liva/pkg/exception"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ScheduleService interface {
	AddSchedule(request model.ScheduleAddRequest) error
	UpdateSchedule(scheduleId string, request model.ScheduleUpdateRequest) error
	FindAll() ([]entity.Schedule, error)
	FindById(scheduleId string) (*entity.Schedule, error)
}
type scheduleService struct {
	scheduleRepository repository.ScheduleRepository
	classRepository    repository.ClassRepository
	subjectRepository  repository.SubjectRepository
	validation         *validator.Validate
	log                *logrus.Logger
}

func NewScheduleService(
	scheduleRepository repository.ScheduleRepository,
	classRepository repository.ClassRepository,
	subjectRepository repository.SubjectRepository,
	validation *validator.Validate,
	log *logrus.Logger,
) ScheduleService {
	return &scheduleService{
		scheduleRepository: scheduleRepository,
		classRepository:    classRepository,
		subjectRepository:  subjectRepository,
		validation:         validation,
		log:                log,
	}
}
func (s *scheduleService) AddSchedule(request model.ScheduleAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countClass, err := s.classRepository.CountById(request.ClassId)
	if err != nil {
		s.log.WithError(err).Error("failed count class to database")
		return err
	}
	countSubject, err := s.subjectRepository.CountById(request.SubjectId)
	if err != nil {
		s.log.WithError(err).Error("failed count subject to database")
		return err
	}
	if countClass < 1 {
		s.log.Warn("class not found")
		return exception.NewError(404, "class not found")
	} else if countSubject < 1 {
		s.log.Warn("subject not found")
		return exception.NewError(404, "subject not found")
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	newTimetable, err := time.ParseInLocation("2006-01-02 15:04:05", request.Timetable, loc)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to date")
		return exception.NewError(400, "invalid timetable format, use 'YYYY-MM-DD HH:MM:SS'")

	}
	schedule := &entity.Schedule{
		ClassId:   request.ClassId,
		SubjectId: request.SubjectId,
		Timetable: newTimetable,
	}
	if err := s.scheduleRepository.Save(*schedule); err != nil {
		s.log.WithError(err).Error("failed save schedule to database")
		return err
	}
	return nil
}
func (s *scheduleService) UpdateSchedule(scheduleId string, request model.ScheduleUpdateRequest) error {
	newScheduleId, err := strconv.ParseInt(scheduleId, 10, 64)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	schedule := new(entity.Schedule)
	countSchedule, err := s.scheduleRepository.CountById(newScheduleId)
	if err != nil {
		s.log.WithError(err).Error("failed count class to database")
		return err
	}
	if countSchedule < 1 {
		s.log.Warn("schedule not found")
		return exception.NewError(404, "schedule not found")
	}
	if request.ClassId != 0 {
		countClass, err := s.classRepository.CountById(request.ClassId)
		if err != nil {
			s.log.WithError(err).Error("failed count class to database")
			return err
		}
		if countClass < 1 {
			s.log.Warn("class not found")
			return exception.NewError(404, "class not found")
		}
		schedule.ClassId = request.ClassId
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
		schedule.SubjectId = request.SubjectId
	}
	if request.Timetable != "" {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		newTimetable, err := time.ParseInLocation("2006-01-02 15:04:05", request.Timetable, loc)
		if err != nil {
			s.log.WithError(err).Error("failed parse string to date")
			return exception.NewError(400, "invalid timetable format, use 'YYYY-MM-DD HH:MM:SS'")
		}
		schedule.Timetable = newTimetable
	}
	if err := s.scheduleRepository.Updates(newScheduleId, *schedule); err != nil {
		s.log.WithError(err).Error("failed save schedule to database")
		return err
	}
	return nil
}
func (s *scheduleService) FindAll() ([]entity.Schedule, error) {
	schedules, err := s.scheduleRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Error("failed find schedules to database")
		return nil, err
	}
	return schedules, nil
}
func (s *scheduleService) FindById(scheduleId string) (*entity.Schedule, error) {
	newScheduleId, err := strconv.ParseInt(scheduleId, 10, 64)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return nil, err
	}
	schedule, err := s.scheduleRepository.FindById(newScheduleId)
	if err != nil {
		s.log.WithError(err).Error("failed find schedule to database")
		return nil, err
	}
	return schedule, nil
}
