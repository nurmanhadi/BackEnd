package repository

import (
	"liva/internal/entity"
	"liva/internal/model"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	Save(teacher entity.Teacher) error
	Updates(teacherId string, request *model.TeacherUpdateRequest) error
	FindAll() ([]entity.Teacher, error)
	FindById(teacherId string) (*entity.Teacher, error)
	FindByNip(nip string) (*entity.Teacher, error)
	CountByNip(nip string) (int64, error)
	CountById(teacherId string) (int64, error)
}
type teacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{db: db}
}
func (r *teacherRepository) Save(teacher entity.Teacher) error {
	return r.db.Save(&teacher).Error
}
func (r *teacherRepository) Updates(teacherId string, request *model.TeacherUpdateRequest) error {
	return r.db.Model(&entity.Teacher{}).Where("id = ?", teacherId).Updates(request).Error
}
func (r *teacherRepository) FindAll() ([]entity.Teacher, error) {
	var teachers []entity.Teacher
	err := r.db.Find(&teachers).Error
	if err != nil {
		return nil, err
	}
	return teachers, nil
}
func (r *teacherRepository) FindById(teacherId string) (*entity.Teacher, error) {
	teacher := new(entity.Teacher)
	err := r.db.Where("id = ?", teacherId).Preload("Subjects").First(&teacher).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}
func (r *teacherRepository) FindByNip(nip string) (*entity.Teacher, error) {
	teacher := new(entity.Teacher)
	err := r.db.Where("nip = ?", nip).First(&teacher).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}
func (r *teacherRepository) CountByNip(nip string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Teacher{}).Where("nip = ?", nip).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *teacherRepository) CountById(teacherId string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Teacher{}).Where("id = ?", teacherId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
