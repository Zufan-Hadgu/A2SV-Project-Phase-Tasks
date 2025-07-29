package Domain

type IJwtService interface {
	GenerateToken(userID string, role string) (string,error)
}