package repository

import (
	"liva/internal/entity"
	"liva/internal/model"

	"gorm.io/gorm"
)

type ClassRepository interface {
	Save(class entity.Class) error
	FindById(classId int) (*entity.Class, error)
	FindAll() ([]entity.Class, error)
	Updates(classId int, request model.ClassUpdateRequest) error
	CountById(classId int) (int64, error)
	CountByTeacherId(teacherId string) (int64, error)
}
type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}
func (r *classRepository) Save(class entity.Class) error {
	return r.db.Save(&class).Error
}
func (r *classRepository) FindById(classId int) (*entity.Class, error) {
	class := new(entity.Class)
	if err := r.db.Where("id = ?", classId).First(&class).Error; err != nil {
		return nil, err
	}
	return class, nil
}
func (r *classRepository) FindAll() ([]entity.Class, error) {
	var classes []entity.Class
	if err := r.db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}
func (r *classRepository) Updates(classId int, request model.ClassUpdateRequest) error {
	return r.db.Model(&entity.Class{}).Where("id = ?", classId).Updates(&request).Error
}
func (r *classRepository) CountById(classId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Class{}).Where("id = ?", classId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *classRepository) CountByTeacherId(teacherId string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Class{}).Where("teacher_id = ?", teacherId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
