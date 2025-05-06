package repository

import (
	"liva/internal/entity"
	"liva/internal/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Save(student entity.Student) error
	FindById(studentId string) (*entity.Student, error)
	FindByNis(nis string) (*entity.Student, error)
	CountByNis(nis string) (int64, error)
	CountById(studentId string) (int64, error)
	CountByEmail(email string) (int64, error)
	FindAll() ([]entity.Student, error)
	Updates(studentId string, request *model.StudentUpdateRequest) error
}
type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}
func (r *studentRepository) Save(student entity.Student) error {
	return r.db.Save(&student).Error
}
func (r *studentRepository) Updates(studentId string, request *model.StudentUpdateRequest) error {
	return r.db.Model(&entity.Student{}).Where("id = ?", studentId).Updates(request).Error
}
func (r *studentRepository) FindAll() ([]entity.Student, error) {
	var students []entity.Student
	err := r.db.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
func (r *studentRepository) FindById(studentId string) (*entity.Student, error) {
	student := new(entity.Student)
	err := r.db.Where("id = ?", studentId).First(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}
func (r *studentRepository) FindByNis(nis string) (*entity.Student, error) {
	student := new(entity.Student)
	err := r.db.Where("nis = ?", nis).First(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}
func (r *studentRepository) CountByNis(nis string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Student{}).Where("nis = ?", nis).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *studentRepository) CountById(studentId string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Student{}).Where("id = ?", studentId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *studentRepository) CountByEmail(email string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Student{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
