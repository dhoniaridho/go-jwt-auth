package users

type CreateUserDto struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UpdateUserDto struct {
	Name     string `form:"name" json:"name" binding:""`
	Email    string `form:"email" json:"email" binding:"email"`
	Password string `form:"password" json:"password" binding:""`
}
