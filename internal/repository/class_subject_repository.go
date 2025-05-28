package repository

import (
	"liva/internal/entity"

	"gorm.io/gorm"
)

type ClassSubjectRepository interface {
	Save(classSubject entity.ClassSubject) error
	Delete(classSubjectId int) error
	CountById(classSubjectId int) (int64, error)
}
type classSubjectRepository struct {
	db *gorm.DB
}

func NewClassSubjectRepository(db *gorm.DB) ClassSubjectRepository {
	return &classSubjectRepository{db: db}
}
func (r *classSubjectRepository) Save(classSubject entity.ClassSubject) error {
	return r.db.Save(&classSubject).Error
}
func (r *classSubjectRepository) Delete(classSubjectId int) error {
	return r.db.Where("id = ?", classSubjectId).Delete(&entity.ClassSubject{}).Error
}
func (r *classSubjectRepository) CountById(classSubjectId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.ClassSubject{}).Where("id = ?", classSubjectId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
