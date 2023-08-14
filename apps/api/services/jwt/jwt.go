package jwt

import (
	"errors"
	"os"
	"time"

	goJwt "github.com/golang-jwt/jwt/v5"
)

var (
	ErrorExpired      = errors.New("token is expired")
	ErrorInvalidToken = errors.New("invalid token")
)

type Payload struct {
	UserId   string
	FullName string
	Email    string
}

type Claims struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	goJwt.RegisteredClaims
}

func EncodeJwt(payload Payload) (string, error) {
	secretKey, ok := os.LookupEnv("JWT_SECRET")

	if !ok {
		panic("JWT secret key is undefined")
	}

	claims := Claims{
		payload.FullName,
		payload.Email,
		goJwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: goJwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  goJwt.NewNumericDate(time.Now()),
			NotBefore: goJwt.NewNumericDate(time.Now()),
			Issuer:    "aurora-auth",
			Subject:   payload.Email,
		},
	}

	token := goJwt.NewWithClaims(goJwt.SigningMethodHS256, claims)
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

	token, err := goJwt.ParseWithClaims(tokenString, &Claims{}, func(token *goJwt.Token) (interface{}, error) {
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
