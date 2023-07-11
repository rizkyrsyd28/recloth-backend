package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rizkyrsyd28/recloth-backend/internal/config"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthUsecase interface {
	Register(c *fiber.Ctx, user model.User) error
	Login(c *fiber.Ctx, username, password string) (fiber.Cookie, error)
	Logout(c *fiber.Ctx) (fiber.Cookie, error)
}

func (u usecase) Register(c *fiber.Ctx, user model.User) error {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	if err := u.repo.CreateUser(c, user); err != nil {
		return err
	}

	return nil
}

func (u usecase) Login(c *fiber.Ctx, username, password string) (fiber.Cookie, error) {

	cookie := fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HTTPOnly: true,
	}

	user, err := u.repo.GetUserByUsername(c, username)
	if err != nil {
		return cookie, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return cookie, err
	}

	claims := &config.JWTClaim{
		UserInfo: model.PublicInfoUser{
			Id:       user.Id,
			Email:    user.Email,
			Username: user.Username,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.Id,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		return cookie, err
	}

	cookie = fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HTTPOnly: true,
	}

	return cookie, err
}

func (u usecase) Logout(c *fiber.Ctx) (cookie fiber.Cookie, err error) {

	cookie = fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}

	return cookie, err
}
