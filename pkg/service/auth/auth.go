package auth

import (
	"errors"
	"github.com/stirk1337/awesomeProject/pkg/repository"
	"github.com/stirk1337/awesomeProject/pkg/user"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user user.User) (int, error) {
	user.Password = hashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	usr, err := s.repo.GetUserHashByUsername(username)
	if err != nil {
		return "", err
	}

	if !checkPasswordHash(password, usr.Password) {
		return "", errors.New("wrong credentials")
	}

	token, err := generateToken(usr)
	if err != nil {
		return "", err
	}

	return token, nil
}
