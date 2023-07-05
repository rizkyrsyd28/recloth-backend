package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rizkyrsyd28/recloth-backend/internal/config"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		tokenRaw := c.Cookies("token")
		if tokenRaw == "" {
			err := errors.New("unauthorized")
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenRaw, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			var _err error
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				_err = errors.New("token invalid")
				return _err
			case jwt.ValidationErrorExpired:
				_err = errors.New("token expired")
				return _err
			default:
				_err = errors.New("token error")
				return _err
			}
		}

		if !token.Valid {
			_err := errors.New("token invalid")
			return _err
		}

		return c.Next()
	}
}
