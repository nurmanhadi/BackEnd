package repository

import (
	"liva/internal/entity"
	"liva/internal/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Save(admin entity.Admin) error
	Updates(adminId string, request *model.AdminUpdateRequest) error
	FindById(adminId string) (*entity.Admin, error)
	FindUsername(username string) (*entity.Admin, error)
	CountById(adminId string) (int64, error)
	CountUsername(username string) (int64, error)
	FindAll() ([]entity.Admin, error)
}
type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}
func (r *adminRepository) Save(admin entity.Admin) error {
	return r.db.Save(&admin).Error
}
func (r *adminRepository) Updates(adminId string, request *model.AdminUpdateRequest) error {
	return r.db.Model(&entity.Admin{}).Where("id = ?", adminId).Updates(request).Error
}
func (r *adminRepository) FindById(adminId string) (*entity.Admin, error) {
	admin := new(entity.Admin)
	err := r.db.Where("id = ?", adminId).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}
func (r *adminRepository) FindUsername(username string) (*entity.Admin, error) {
	admin := new(entity.Admin)
	err := r.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}
func (r *adminRepository) CountById(adminId string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Teacher{}).Where("id = ?", adminId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *adminRepository) CountUsername(username string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Teacher{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *adminRepository) FindAll() ([]entity.Admin, error) {
	var admins []entity.Admin
	err := r.db.Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}
