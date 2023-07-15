package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/helper"
	"github.com/rizkyrsyd28/recloth-backend/internal/usecase"
)

func GetCart(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		payload, err := helper.ParseJWTPayload(c)
		if err != nil {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		data, err := u.GetCart(c, payload.Id)
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

/*
	{
		"method" 		: "add",
		"itemId"		: item_id,
	}
*/

func UpdateCart(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		payload, err := helper.ParseJWTPayload(c)
		if err != nil {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		update := struct {
			Method string `json:"method"`
			ItemId string `json:"item_id"`
		}{}

		if err := json.Unmarshal(c.Body(), &update); err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		if update.Method == "clear" {

			err = u.ClearCart(c, payload.Id)
			if err != nil {

				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": err.Error(),
				})
				return nil

			}

		} else {

			err = u.UpdateCart(c, update.Method, update.ItemId, payload.Id)
			if err != nil {

				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": err.Error(),
				})
				return nil

			}

		}

		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "updated !",
		})

		return nil
	}
}
