package config

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
)

var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
