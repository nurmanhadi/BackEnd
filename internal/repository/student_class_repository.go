package repository

import "gorm.io/gorm"

type StudentClassRepository interface{}
type studentClassRepository struct {
	db *gorm.DB
}

func NewStudentClassRepository(db *gorm.DB) StudentClassRepository {
	return &studentClassRepository{db: db}
}
