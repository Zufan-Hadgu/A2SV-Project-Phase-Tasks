package Usecases_test

import (
	"errors"
	"task_manager/Domain"
	"task_manager/Usecases"
	"task_manager/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister_UserAlreadyExists(t *testing.T) {
	// Arrange
	mockUserRepo := new(mocks.MockUserRepository)
	mockPass := new(mocks.MockPasswordService)
	mockJWT := new(mocks.MockJwtService)

	usecase := Usecases.NewUserUsecase(mockUserRepo, mockJWT, mockPass)

	existingUser := Domain.User{
		Username: "zufan",
		Password: "hashedpassword",
	}

	
	mockUserRepo.On("FindByUsername", "zufan").Return(&existingUser, nil)

	// Act
	err := usecase.Register(Domain.User{
		Username: "zufan",
		Password: "somepass",
	})

	// Assert
	assert.EqualError(t, err, "user already exist")
	mockUserRepo.AssertExpectations(t)
}

func TestRegister_Success(t *testing.T) {
	// Arrange
	mockUserRepo := new(mocks.MockUserRepository)
	mockPass := new(mocks.MockPasswordService)
	mockJWT := new(mocks.MockJwtService)

	usecase := Usecases.NewUserUsecase(mockUserRepo, mockJWT, mockPass)

	newUser := Domain.User{
		Username: "newuser",
		Password: "plaintext",
	}

	
	mockUserRepo.On("FindByUsername", "newuser").Return(nil, errors.New("not found"))
	mockPass.On("HashPassword", "plaintext").Return("hashed", nil)
	mockUserRepo.On("CountDB").Return(int64(0), nil)
	mockUserRepo.On("Create", Domain.User{Username: "newuser", Password: "hashed", Role: "Admin"}).Return(nil)

	err := usecase.Register(newUser)

	
	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
	mockPass.AssertExpectations(t)
}

func TestLogin_InvalidPassword(t *testing.T) {
	// Arrange
	mockUserRepo := new(mocks.MockUserRepository)
	mockPass := new(mocks.MockPasswordService)
	mockJWT := new(mocks.MockJwtService)

	usecase := Usecases.NewUserUsecase(mockUserRepo, mockJWT, mockPass)

	user := Domain.User{
		Username: "zufan",
		Password: "hashedpassword",
		Role:     "user",
	}

	
	mockUserRepo.On("FindByUsername", "zufan").Return(&user, nil)
	mockPass.On("ComparePassword", "hashedpassword", "wrongpass").Return(errors.New("invalid"))

	// Act
	_, err := usecase.Login(Domain.LoginRequest{
		Username: "zufan",
		Password: "wrongpass",
	})

	// Assert
	assert.EqualError(t, err, "invalid password")
	mockUserRepo.AssertExpectations(t)
	mockPass.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	// Arrange
	mockUserRepo := new(mocks.MockUserRepository)
	mockPass := new(mocks.MockPasswordService)
	mockJWT := new(mocks.MockJwtService)

	usecase := Usecases.NewUserUsecase(mockUserRepo, mockJWT, mockPass)

	user := Domain.User{
		Username: "zufan",
		Password: "hashedpassword",
		Role:     "user",
	}

	mockUserRepo.On("FindByUsername", "zufan").Return(&user, nil)
	mockPass.On("ComparePassword", "hashedpassword", "correctpass").Return(nil)
	mockJWT.On("GenerateToken", "zufan", "user").Return("token123", nil)

	// Act
	token, err := usecase.Login(Domain.LoginRequest{
		Username: "zufan",
		Password: "correctpass",
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "token123", token)
	mockUserRepo.AssertExpectations(t)
	mockPass.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}
