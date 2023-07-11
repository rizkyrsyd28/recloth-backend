package config

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
	"os"
)

var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	UserInfo model.PublicInfoUser
	jwt.RegisteredClaims
}
