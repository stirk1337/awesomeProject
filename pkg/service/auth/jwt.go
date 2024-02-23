package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stirk1337/awesomeProject/pkg/user"
	"os"
	"time"
)

const (
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func generateToken(usr user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		usr.Id,
	})
	return token.SignedString([]byte(os.Getenv("JWT_TOKEN")))
}
