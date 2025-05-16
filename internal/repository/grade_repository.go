package repository

import (
	"liva/internal/entity"
	"liva/internal/model"

	"gorm.io/gorm"
)

type GradeRepository interface {
	Save(grade *entity.Grade) error
	Updates(gradeId int64, request *model.GradeUpdateRequest) error
	CountById(gradeId int64) (int64, error)
	FindAll() ([]entity.Grade, error)
	FindById(gradeId int64) (*entity.Grade, error)
}
type gradeRepository struct {
	db *gorm.DB
}

func NewGradeRepository(db *gorm.DB) GradeRepository {
	return &gradeRepository{db: db}
}

func (r *gradeRepository) Save(grade *entity.Grade) error {
	return r.db.Save(&grade).Error
}
func (r *gradeRepository) Updates(gradeId int64, request *model.GradeUpdateRequest) error {
	return r.db.Model(&entity.Grade{}).Where("id = ?", gradeId).Updates(&request).Error
}
func (r *gradeRepository) CountById(gradeId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Grade{}).Where("id = ?", gradeId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *gradeRepository) FindAll() ([]entity.Grade, error) {
	var grades []entity.Grade
	err := r.db.Find(&grades).Error
	if err != nil {
		return nil, err
	}
	return grades, nil
}
func (r *gradeRepository) FindById(gradeId int64) (*entity.Grade, error) {
	grade := new(entity.Grade)
	err := r.db.Where("id = ?", gradeId).First(&grade).Error
	if err != nil {
		return nil, err
	}
	return grade, nil
}
