package auth

import (
	"api/src/internal/apps/users"
	"api/src/internal/pkg/jwt"
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) SignIn(payload *LoginDto) (gin.H, error) {

	user, err := users.UserService.FindOneByEmail(users.UserService{}, payload.Email)

	if err != nil {
		return gin.H{
			"token": nil,
			"user":  nil,
		}, err
	}

	validateErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if validateErr != nil {
		return gin.H{
			"token": nil,
			"user":  nil,
		}, errors.New("password mismatch")
	}

	token, jwtErr := jwt.GenerateToken(jwt.Payload{
		Uid: user.Uid,
	})

	if jwtErr != nil {
		return gin.H{}, jwtErr
	}

	if err != nil {
		return gin.H{}, err
	}

	u := gin.H{
		"token": token,
		"user":  user,
	}

	return u, err

}

func (s *AuthService) Register(payload *RegisterDto) (gin.H, error) {
	user, err := users.UserService.CreateOne(users.UserService{}, &users.CreateUserDto{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		return gin.H{}, err
	}

	token, jwtErr := jwt.GenerateToken(jwt.Payload{
		Uid: user.Uid,
	})

	if jwtErr != nil {
		return gin.H{}, jwtErr
	}

	return gin.H{
		"user":  user,
		"token": token,
	}, nil
}
