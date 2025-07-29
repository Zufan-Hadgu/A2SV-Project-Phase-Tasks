package Domain

type IPasswordService interface {
	HashPassword (password string) (string,error)
	ComparePassword(hashedPassword,password string,) error 
}