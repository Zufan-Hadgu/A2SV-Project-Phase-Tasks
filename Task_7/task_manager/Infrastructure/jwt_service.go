package Infrastructure

import (
	"errors"
	"task_manager/Domain"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("mysecret")

func JwtService(user Domain.User, inputPassword string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),  
	})

	token, err := tokenString.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("internal server error")
	}

	return token, nil
}