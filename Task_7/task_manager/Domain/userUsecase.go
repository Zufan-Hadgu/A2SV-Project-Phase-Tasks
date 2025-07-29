package Domain



type IUserUsecase interface {
	Register(user User)(string,error)
	Login(loginrequest LoginRequest)(string,error)
}