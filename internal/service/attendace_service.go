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

type AttendanceService interface {
	AddAttendance(request model.AttendanceAddRequest) error
	UpdateAttendance(attendaceId string, request model.AttendanceUpdateRequest) error
	FindAll() ([]entity.Attendance, error)
	FindById(attendaceId string) (*entity.Attendance, error)
}
type attendanceService struct {
	attendanceRepository repository.AttendanceRepository
	studentRepository    repository.StudentRepository
	scheduleRepository   repository.ScheduleRepository
	validation           *validator.Validate
	log                  *logrus.Logger
}

func NewAttendaceService(
	attendanceRepository repository.AttendanceRepository,
	studentRepository repository.StudentRepository,
	scheduleRepository repository.ScheduleRepository,
	validation *validator.Validate,
	log *logrus.Logger,
) AttendanceService {
	return &attendanceService{
		attendanceRepository: attendanceRepository,
		studentRepository:    studentRepository,
		scheduleRepository:   scheduleRepository,
		validation:           validation,
		log:                  log,
	}
}
func (s *attendanceService) AddAttendance(request model.AttendanceAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countSchedule, err := s.scheduleRepository.CountById(request.ScheduleId)
	if err != nil {
		s.log.WithError(err).Error("failed count schedule to database")
		return err
	}
	countStudent, err := s.studentRepository.CountById(request.StudentId)
	if err != nil {
		s.log.WithError(err).Error("failed count student to database")
		return err
	}
	if countSchedule < 1 {
		s.log.Warn("schedule not found")
		return exception.NewError(404, "schedule not found")
	} else if countStudent < 1 {
		s.log.Warn("student not found")
		return exception.NewError(404, "student not found")
	}
	attendance := &entity.Attendance{
		StudentId:  request.StudentId,
		ScheduleId: request.ScheduleId,
		Status:     request.Status,
	}
	if err := s.attendanceRepository.Save(*attendance); err != nil {
		s.log.WithError(err).Error("failed save attendace to database")
		return err
	}
	return nil
}
func (s *attendanceService) UpdateAttendance(attendaceId string, request model.AttendanceUpdateRequest) error {
	newAttendanceId, err := strconv.ParseInt(attendaceId, 10, 64)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countAttendace, err := s.attendanceRepository.CountById(newAttendanceId)
	if err != nil {
		s.log.WithError(err).Error("failed count attendace to database")
		return err
	}
	countSchedule, err := s.scheduleRepository.CountById(request.ScheduleId)
	if err != nil {
		s.log.WithError(err).Error("failed count schedule to database")
		return err
	}
	countStudent, err := s.studentRepository.CountById(request.StudentId)
	if err != nil {
		s.log.WithError(err).Error("failed count student to database")
		return err
	}
	if countSchedule < 1 {
		s.log.Warn("schedule not found")
		return exception.NewError(404, "schedule not found")
	} else if countStudent < 1 {
		s.log.Warn("student not found")
		return exception.NewError(404, "student not found")
	} else if countAttendace < 1 {
		s.log.Warn("attendance not found")
		return exception.NewError(404, "attendace not found")
	}
	if err := s.attendanceRepository.Updates(newAttendanceId, request); err != nil {
		s.log.WithError(err).Error("failed update attendace to database")
		return err
	}
	return nil
}
func (s *attendanceService) FindAll() ([]entity.Attendance, error) {
	attendaces, err := s.attendanceRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Error("failed find attendaces to database")
		return nil, err
	}
	return attendaces, nil
}
func (s *attendanceService) FindById(attendaceId string) (*entity.Attendance, error) {
	newAttendanceId, err := strconv.ParseInt(attendaceId, 10, 64)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return nil, err
	}
	attendace, err := s.attendanceRepository.FindById(newAttendanceId)
	if err != nil {
		s.log.WithError(err).Warn("failed find attendace to database")
		return nil, err
	}
	return attendace, nil
}
