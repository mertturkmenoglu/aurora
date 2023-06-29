package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var (
	ErrorExpired      = errors.New("token is expired")
	ErrorInvalidToken = errors.New("invalid token")
)

type Payload struct {
	Id       string
	FullName string
	Email    string
}

type Claims struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func EncodeJwt(payload Payload) (string, error) {
	secretKey, ok := os.LookupEnv("JWT_SECRET")

	if !ok {
		panic("JWT secret key is undefined")
	}

	claims := Claims{
		payload.FullName,
		payload.Email,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "aurora-auth",
			Subject:   payload.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return signed, nil
}

func DecodeJwt(tokenString string) (*Claims, error) {
	secretKey, ok := os.LookupEnv("JWT_SECRET")

	if !ok {
		panic("JWT secret key is undefined")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, ErrorInvalidToken
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, ErrorExpired
	}

	return claims, nil
}
