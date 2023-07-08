package service

import (
	"RESTful/app/repositories"
	"RESTful/internal/domain"
	"RESTful/pkg/helpers"
	"errors"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (as *AuthService) Authenticate(username, password string) (*domain.User, error) {
	user, err := as.userRepository.GetByName(username)
	if err != nil {
		return nil, err
	}

	// Compare password
	if !helpers.ComparePassword(user.Password, password) {
		return nil, errors.New("invalid credentials because password is wrong")
	}

	return user, nil
}
