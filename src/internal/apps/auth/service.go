package auth

import (
	"api/src/config/env"
	"api/src/internal/apps/users"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "web",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"sub": "web",
		"id":  user.ID,
	})

	secretKey := []byte(env.Get("JWT_SECRET_KEY"))

	tokenString, err := token.SignedString(secretKey)

	u := gin.H{
		"token": tokenString,
		"user":  user,
	}

	if err != nil {
		return gin.H{
			"token": tokenString,
			"user":  user,
		}, err
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "web",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"sub": "web",
		"id":  user.Uid,
	})

	secretKey := []byte(env.Get("JWT_SECRET_KEY"))

	tokenString, jwtErr := token.SignedString(secretKey)

	if jwtErr != nil {
		return gin.H{"user": user}, err
	}

	return gin.H{
		"user":  user,
		"token": tokenString,
	}, nil
}
