package services

import (
	"auto-daily-report/src/config"
	auth "auto-daily-report/src/repositories/authentication"
)

type AuthServiceInterface interface {
}

// AuthService is the implementation of the UserService interface
type AuthService struct {
	config         *config.Config
	userRepository auth.AuthRepositoryInterface
}

// NewUserService creates a new UserService
func NewAuthService(config *config.Config, userRepository auth.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		config:         config,
		userRepository: userRepository,
	}
}
