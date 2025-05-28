package repository

import (
	"liva/internal/entity"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Save(schedule entity.Schedule) error
	CountById(scheduleId int64) (int64, error)
	Updates(scheduleId int64, schedule entity.Schedule) error
	FindAll() ([]entity.Schedule, error)
	FindById(scheduleId int64) (*entity.Schedule, error)
}
type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}
func (r *scheduleRepository) Save(schedule entity.Schedule) error {
	return r.db.Save(&schedule).Error
}
func (r *scheduleRepository) Updates(scheduleId int64, schedule entity.Schedule) error {
	return r.db.Model(&entity.Schedule{}).Where("id = ?", scheduleId).Updates(&schedule).Error
}
func (r *scheduleRepository) CountById(scheduleId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subject{}).Where("id = ?", scheduleId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *scheduleRepository) FindAll() ([]entity.Schedule, error) {
	var schedules []entity.Schedule
	err := r.db.Find(&schedules).Error
	if err != nil {
		return nil, err
	}
	return schedules, nil
}
func (r *scheduleRepository) FindById(scheduleId int64) (*entity.Schedule, error) {
	schedule := new(entity.Schedule)
	err := r.db.Where("id = ?", scheduleId).Preload("Attendaces").First(&schedule).Error
	if err != nil {
		return nil, err
	}
	return schedule, nil
}
