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
	UserLogin(request *model.AuthLoginRequest) (*model.AuthResponse, error)
}
type authService struct {
	userRepository repository.UserRepository
	validation     *validator.Validate
	log            *logrus.Logger
	viper          *viper.Viper
}

func NewAuthService(
	userRepository repository.UserRepository,
	validation *validator.Validate,
	log *logrus.Logger,
	viper *viper.Viper,
) AuthService {
	return &authService{
		userRepository: userRepository,
		validation:     validation,
		log:            log,
		viper:          viper,
	}
}
func (s *authService) UserLogin(request *model.AuthLoginRequest) (*model.AuthResponse, error) {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation request")
		return nil, err
	}
	user, err := s.userRepository.FindByIdentifier(request.Identifier)
	if err != nil {
		s.log.WithError(err).Warn("failed find user from database")
		return nil, exception.NewError(400, "identifier or password wrong")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		s.log.WithError(err).Warn("failed compare hash password")
		return nil, exception.NewError(400, "identifier or password wrong")
	}
	token, err := JwtGenerateToken(user.Id, user.ReferenceId, user.Role, []byte(s.viper.GetString("jwt.key")))
	if err != nil {
		s.log.WithError(err).Error("failed generate token jwt")
		return nil, err
	}

	return &model.AuthResponse{
		Token: token,
		Role:  user.Role,
	}, nil
}
