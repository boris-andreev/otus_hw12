package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	secretKey               = []byte("go-gin-todo-list-api")
	errInvalidSigningMethod = errors.New("Invalid signing method")
	errInvalidToken         = errors.New("Invalid token")
)

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "go-gin-todo-list-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidSigningMethod
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errInvalidToken
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}

	return nil, errInvalidToken
}
