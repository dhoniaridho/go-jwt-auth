package auth

type LoginDto struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type RegisterDto struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
