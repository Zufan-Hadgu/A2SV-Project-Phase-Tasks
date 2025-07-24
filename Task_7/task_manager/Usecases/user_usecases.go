package Usecases

import (
	"context"
	"task_manager/Domain"
)

type UserUsecase struct {
	UserRepo Domain.UserRepository
}


func NewUserUsecase(repo Domain.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

func (uc *UserUsecase) Register(ctx context.Context,user Domain.User)error {
	return uc.UserRepo.Register(ctx,user)

}

func (uc *UserUsecase) Login(ctx context.Context,loginRequest Domain.LoginRequest) (string,error){
	return uc.UserRepo.Login(ctx,loginRequest)
}



