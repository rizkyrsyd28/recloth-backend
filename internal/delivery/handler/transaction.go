package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/helper"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
	"github.com/rizkyrsyd28/recloth-backend/internal/usecase"
)

func Checkout(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var transaction model.Transaction

		payload, err := helper.ParseJWTPayload(c)
		if err != nil {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		if err := json.Unmarshal(c.Body(), &transaction); err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		transaction.UserId = payload.Id

		if err := u.Checkout(c, transaction); err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "transaction success !",
		})

		return nil
	}
}
func GetTransaction(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		payload, err := helper.ParseJWTPayload(c)
		if err != nil {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		userId := payload.Id

		data, err := u.GetTransactionByUserId(c, userId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": data,
		})

		return nil
	}
}
