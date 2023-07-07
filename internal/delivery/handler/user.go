package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
	"github.com/rizkyrsyd28/recloth-backend/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var userInput model.User

		if err := json.Unmarshal(c.Body(), &userInput); err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		if err := u.Register(c, userInput); err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "register success !",
		})

		return nil
	}
}

func Login(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var userInput model.User

		if err := json.Unmarshal(c.Body(), &userInput); err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		cookie, err := u.Login(c, userInput.Username, userInput.Password)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "username atau password tidak tepat",
				})
				return nil
			} else {
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"messa ge": "password salah",
				})
				return nil
			}
		}

		c.Cookie(&cookie)
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "authorized !",
		})

		return nil
	}
}

func Logout(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		cookie, err := u.Logout(c)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		c.Cookie(&cookie)
		c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": "unauthorized",
		})

		return nil
	}
}
