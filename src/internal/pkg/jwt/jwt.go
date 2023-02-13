package jwt

import (
	"api/src/config/env"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(env.Get("JWT_SECRET_KEY"))

type Payload struct {
	Uid string
}

func ValidateToken(tokenString string) error {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("failed to parse token:", err)
		return err
	}

	// Validate the token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check the token expiration
		exp, ok := claims["exp"].(float64)

		if ok {
			return nil
		}

		if !ok {
			return errors.New("failed to retrieve token expiration")
		}
		expiration := time.Unix(int64(exp), 0)
		if time.Now().After(expiration) {
			return errors.New("token has expired")
		}

	} else {
		return errors.New("token is not valid")
	}
	return nil

}

func GenerateToken(payload Payload) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "web",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"sub": "web",
		"uid": payload.Uid,
	})

	tokenString, jwtErr := token.SignedString(secretKey)

	if jwtErr != nil {
		return "", jwtErr
	}

	return tokenString, nil

}
