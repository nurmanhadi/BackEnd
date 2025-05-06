package service

import (
	"liva/internal/model"
	"liva/internal/repository"
	"liva/pkg/exception"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	StudentLogin(request *model.StudentLoginRequest) (*model.TokenResponse, error)
}
type authService struct {
	studentRepository repository.StudentRepository
	validation        *validator.Validate
	log               *logrus.Logger
	viper             *viper.Viper
}

func NewAuthService(
	studentRepository repository.StudentRepository,
	validation *validator.Validate,
	log *logrus.Logger,
	viper *viper.Viper,
) AuthService {
	return &authService{
		studentRepository: studentRepository,
		validation:        validation,
		log:               log,
		viper:             viper,
	}
}
func (s *authService) StudentLogin(request *model.StudentLoginRequest) (*model.TokenResponse, error) {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation request")
		return nil, err
	}
	student, err := s.studentRepository.FindByNis(request.Nis)
	if err != nil {
		s.log.WithError(err).Warn("failed find student from database")
		return nil, exception.NewError(400, "nis or password wrong")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(request.Password)); err != nil {
		s.log.WithError(err).Warn("failed find student from database")
		return nil, exception.NewError(400, "nis or password wrong")
	}
	token, err := JwtGenerateToken(student.Id, []byte(s.viper.GetString("jwt.key")))
	if err != nil {
		s.log.WithError(err).Error("failed generate token jwt")
		return nil, err
	}
	response := &model.TokenResponse{
		Token: token,
	}
	return response, nil
}
