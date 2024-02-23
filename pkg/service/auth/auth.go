package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/stirk1337/awesomeProject/pkg/repository"
	"github.com/stirk1337/awesomeProject/pkg/user"
	"os"
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

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	return claims.UserId, nil
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
