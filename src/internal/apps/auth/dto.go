package auth

type LoginDto struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
