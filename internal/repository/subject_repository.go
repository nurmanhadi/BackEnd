package repository

import (
	"liva/internal/entity"
	"liva/internal/model"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	Save(subject entity.Subject) error
	Updates(subjectId int, request model.SubjectUpdateRequest) error
	FindById(subjectId int) (*entity.Subject, error)
	CountById(subjectId int) (int64, error)
	FindAll() ([]entity.Subject, error)
}
type subjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{db: db}
}
func (r *subjectRepository) Save(subject entity.Subject) error {
	return r.db.Save(&subject).Error
}
func (r *subjectRepository) Updates(subjectId int, request model.SubjectUpdateRequest) error {
	return r.db.Model(&entity.Subject{}).Where("id = ?", subjectId).Updates(&request).Error
}
func (r *subjectRepository) FindById(subjectId int) (*entity.Subject, error) {
	subject := new(entity.Subject)
	err := r.db.Where("id = ?", subjectId).First(&subject).Error
	if err != nil {
		return nil, err
	}
	return subject, nil
}
func (r *subjectRepository) CountById(subjectId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subject{}).Where("id = ?", subjectId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *subjectRepository) FindAll() ([]entity.Subject, error) {
	var subjects []entity.Subject
	err := r.db.Find(&subjects).Error
	if err != nil {
		return nil, err
	}
	return subjects, nil
}
