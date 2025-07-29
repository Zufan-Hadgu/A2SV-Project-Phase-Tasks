package Infrastructure

import (
"golang.org/x/crypto/bcrypt"
)

type passwordService struct{}

func NewPasswordService() *passwordService {
	return &passwordService{}
}

func (ps *passwordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "",err
	}
	return string(hashedPassword),nil

}

func (ps *passwordService) ComparePassword (hashedPassword,password string) error{

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))

}