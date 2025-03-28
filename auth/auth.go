package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// authentication logic
var jwtKey = []byte("supersecretkey")

type Claims struct {
	DeviceID         string `json:"device_id"`
	RegisteredClaims jwt.RegisteredClaims
}

func (c *Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}

func GenerateJWT(deviceID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		DeviceID: deviceID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
