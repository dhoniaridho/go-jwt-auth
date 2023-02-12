package auth

import (
	"api/src/internal/apps/users"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) SignIn(payload *LoginDto) (string, error) {

	user, err := users.UserService.FindOneByEmail(users.UserService{}, payload.Email)

	if err != nil {
		return "", err
	}

	validateErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if validateErr != nil {
		return "", errors.New("password mismatch")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "web",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"sub": "web",
		"id":  user.ID,
	})

	secretKey := []byte("my-secret-key")

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, err

}
