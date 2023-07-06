package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
	"github.com/rizkyrsyd28/recloth-backend/internal/usecase"
	"strconv"
)

func GetItems(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		page, err := strconv.Atoi(c.Params("page"))
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		data, err := u.GetAllItems(c, page, limit)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": data,
		})

		return nil
	}
}

func GetItem(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")

		data, err := u.GetItemById(c, id)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": data,
		})

		return nil
	}
}

func UpdateItem(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")

		var request struct {
			Key   string      `json:"key"`
			Value interface{} `json:"value"`
		}

		if err := json.Unmarshal(c.Body(), &request); err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		if err := u.UpdateItemById(c, id, request.Key, request.Value); err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		c.Status(fiber.StatusCreated).JSON(fiber.Map{})

		return nil
	}
}

func DeleteItem(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")

		if err := u.DeleteItemById(c, id); err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		c.Status(fiber.StatusNoContent).JSON(fiber.Map{})

		return nil
	}
}

func PostItem(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var inputItem model.Item
		if err := json.Unmarshal(c.Body(), &inputItem); err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		if err := u.PostItem(c, inputItem); err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return err
		}

		c.Status(fiber.StatusCreated).JSON(fiber.Map{})

		return nil
	}
}

func Test(u usecase.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "safe",
		})

		return nil
	}
}
