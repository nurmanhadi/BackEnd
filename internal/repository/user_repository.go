package repository

import (
	"liva/internal/entity"
	"liva/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByIdentifier(identifier string) (*entity.User, error)
	Save(user *entity.User) error
	Updates(referenceId string, request model.UserUpdateRequest) error
	CountByReferenceId(referenceId string) (int64, error)
	CountByIdentifier(identifier string) (int64, error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) FindByIdentifier(identifier string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.Where("identifier = ?", identifier).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userRepository) Save(user *entity.User) error {
	return r.db.Save(&user).Error
}
func (r *userRepository) Updates(referenceId string, request model.UserUpdateRequest) error {
	return r.db.Model(&entity.User{}).Where("reference_id = ?", referenceId).Updates(&request).Error
}
func (r *userRepository) CountByReferenceId(referenceId string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Where("reference_id = ?", referenceId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *userRepository) CountByIdentifier(identifier string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Where("identifier = ?", identifier).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
