package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/log/v8"
)

var jwtKey = []byte("dappers_lab_secret_key")

type JwtClaims struct {
	jwt.StandardClaims
	Email string
}

func Generate(email string) (string, error) {
	claims := &JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
		Email: email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Verify(jwtToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		log.Debug("Error while parsing token")
		return nil, errors.New("Couldn't parse claims")

	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT token expired")

	}
	return claims, nil
}
