package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rizkyrsyd28/recloth-backend/internal/config"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
)

func ParseJWTPayload(c *fiber.Ctx) (model.PublicInfoUser, error) {

	tokenRaw := c.Cookies("token")

	claims := &config.JWTClaim{}

	_, err := jwt.ParseWithClaims(tokenRaw, claims, func(t *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})
	if err != nil {
		return model.PublicInfoUser{}, err
	}

	return claims.UserInfo, nil
}
