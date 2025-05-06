package repository

import (
	"liva/internal/entity"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Save(student entity.Student) error
	FindById(studentId string) (*entity.Student, error)
	FindByNis(nis string) (*entity.Student, error)
	CountByNis(nis string) (int64, error)
	CountById(studentId string) (int64, error)
	CountByEmail(email string) (int64, error)
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
