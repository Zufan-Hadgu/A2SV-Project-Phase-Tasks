package Infrastructure

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("mysecret")
type jwtService struct{}

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (jw *jwtService) GenerateToken(username,role string) (string, error) {
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role": role,
	})
	token, err := tokenString.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("internal server error")
	}

	return token, nil
}
