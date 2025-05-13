package repository

import (
	"liva/internal/entity"
	"liva/internal/model"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Save(attendance entity.Attendance) error
	CountById(attendaceId int64) (int64, error)
	Updates(attendaceId int64, request model.AttendanceUpdateRequest) error
	FindAll() ([]entity.Attendance, error)
	FindById(attendaceId int64) (*entity.Attendance, error)
}
type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendaceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}
func (r *attendanceRepository) Save(attendance entity.Attendance) error {
	return r.db.Save(&attendance).Error
}
func (r *attendanceRepository) CountById(attendaceId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Attendance{}).Where("id = ?", attendaceId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *attendanceRepository) Updates(attendaceId int64, request model.AttendanceUpdateRequest) error {
	return r.db.Model(&entity.Attendance{}).Where("id = ?", attendaceId).Updates(&request).Error
}
func (r *attendanceRepository) FindAll() ([]entity.Attendance, error) {
	var attendaces []entity.Attendance
	err := r.db.Find(&attendaces).Error
	if err != nil {
		return nil, err
	}
	return attendaces, nil
}
func (r *attendanceRepository) FindById(attendaceId int64) (*entity.Attendance, error) {
	attendace := new(entity.Attendance)
	err := r.db.Where("id = ?", attendaceId).First(&attendace).Error
	if err != nil {
		return nil, err
	}
	return attendace, nil
}
