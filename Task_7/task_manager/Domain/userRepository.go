package Domain



type IUserRepository interface {
	Create(user User) error
	FindByUsername(username string)(*User,error) 
	CountDB()(int64,error)
}