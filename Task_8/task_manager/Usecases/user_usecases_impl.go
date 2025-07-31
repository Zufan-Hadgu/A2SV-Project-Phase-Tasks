package Usecases

import (

	"errors"
	"task_manager/Domain"
	

)

type UserUsecase struct {
	UserRepo Domain.IUserRepository
	jwtSer Domain.IJwtService
	pasSer Domain.IPasswordService

}


func NewUserUsecase(repo Domain.IUserRepository,jwtserv Domain.IJwtService, passServ Domain.IPasswordService) *UserUsecase {
	return &UserUsecase{
		UserRepo: repo,
		jwtSer : jwtserv,
		pasSer : passServ,

	}
}

func (uc *UserUsecase) Register(user Domain.User)error {

	_,err := uc.UserRepo.FindByUsername(user.Username)
	if err == nil{
		return errors.New("user already exist")
	}
	hashedPassword, err := uc.pasSer.HashPassword(user.Password)
	if err != nil{
		return errors.New("Internal Error")
		
	}
	user.Password = hashedPassword
	count,err := uc.UserRepo.CountDB()
	user.Role = "user"
	if count == 0{
		user.Role = "Admin"
	}
	err = uc.UserRepo.Create(user)
	if err != nil{
		return errors.New("Internal Error")
	}
	return nil

}
func (uc *UserUsecase) Login(loginRequest Domain.LoginRequest) (string, error) {
	user, err := uc.UserRepo.FindByUsername(loginRequest.Username)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = uc.pasSer.ComparePassword(user.Password, loginRequest.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}


	// Generate token
	token, err := uc.jwtSer.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}




