package service

import (
	"liva/internal/entity"
	"liva/internal/model"
	"liva/internal/repository"
	"liva/pkg"
	"liva/pkg/exception"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	AddAdmin(request model.AdminAddRequest) error
	UpdateAdmin(adminId string, request model.AdminUpdateRequest) error
	FindById(adminId string) (*entity.Admin, error)
	FindAll() ([]entity.Admin, error)
}
type adminService struct {
	adminRepository repository.AdminRepository
	userRepository  repository.UserRepository
	validation      *validator.Validate
	log             *logrus.Logger
}

func NewAdminService(adminRepository repository.AdminRepository,
	userRepository repository.UserRepository,
	validation *validator.Validate,
	log *logrus.Logger) AdminService {
	return &adminService{
		adminRepository: adminRepository,
		userRepository:  userRepository,
		validation:      validation,
		log:             log,
	}
}
func (s *adminService) AddAdmin(request model.AdminAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countUsername, err := s.adminRepository.CountUsername(request.Username)
	if err != nil {
		s.log.WithError(err).Error("failed count username from database")
		return err
	}
	if countUsername > 0 {
		s.log.Warn("username already exists")
		return exception.NewError(409, "username already exists")
	}
	adminId := uuid.NewString()
	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed hash password from bcrypt")
		return err
	}
	countIdentifier, err := s.userRepository.CountByIdentifier(request.Username)
	if err != nil {
		s.log.WithError(err).Error("failed count identifier from database")
		return err
	}
	if countIdentifier > 0 {
		s.log.Warn("identifier already exists")
		return exception.NewError(409, "identifier already exists")
	}
	countReference, err := s.userRepository.CountByReferenceId(adminId)
	if err != nil {
		s.log.WithError(err).Error("failed count reference_id from database")
		return err
	}
	if countReference > 0 {
		s.log.Warn("reference_id already exists")
		return exception.NewError(409, "reference_id already exists")
	}
	if err := s.userRepository.Save(&entity.User{
		Id:          uuid.NewString(),
		Identifier:  request.Username,
		Password:    string(newPassword),
		Role:        string(pkg.RoleAdmin),
		ReferenceId: adminId,
	}); err != nil {
		s.log.WithError(err).Error("failed save user to database")
		return err
	}
	admin := &entity.Admin{
		Id:       adminId,
		Username: request.Username,
	}
	err = s.adminRepository.Save(*admin)
	if err != nil {
		s.log.WithError(err).Error("failed save admin to database")
		return err
	}
	return nil
}
func (s *adminService) UpdateAdmin(adminId string, request model.AdminUpdateRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countId, err := s.adminRepository.CountById(adminId)
	if err != nil {
		s.log.WithError(err).Error("failed count username from database")
		return err
	}
	if countId < 1 {
		s.log.Warn("admin not found")
		return exception.NewError(404, "admin not found")
	}

	err = s.adminRepository.Updates(adminId, &request)
	if err != nil {
		s.log.WithError(err).Error("failed update admin to database")
		return err
	}
	return nil
}
func (s *adminService) FindById(adminId string) (*entity.Admin, error) {
	admin, err := s.adminRepository.FindById(adminId)
	if err != nil {
		s.log.WithError(err).Warn("admin not found")
		return nil, exception.NewError(404, "admin not found")
	}
	return admin, nil
}
func (s *adminService) FindAll() ([]entity.Admin, error) {
	admins, err := s.adminRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Warn("admins not found")
		return nil, exception.NewError(404, "admins not found")
	}
	return admins, nil
}
