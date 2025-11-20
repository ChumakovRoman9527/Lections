package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(email string) (string, error) {
	// jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	claims := jwt.MapClaims{"email": email}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return s, nil
}
