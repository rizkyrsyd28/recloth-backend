package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
)

type ItemUsecase interface {
	GetAllItems(c *fiber.Ctx, page, limit int) ([]model.Item, error)
	GetItemById(c *fiber.Ctx, id string) (model.Item, error)
	UpdateItemById(c *fiber.Ctx, id, attribute string, value interface{}) error
	DeleteItemById(c *fiber.Ctx, id string) error
	PostItem(c *fiber.Ctx, item model.Item) error
}

func (u usecase) GetAllItems(c *fiber.Ctx, page, limit int) ([]model.Item, error) {
	return u.repo.GetAllItems(c, page, limit)
}

func (u usecase) GetItemById(c *fiber.Ctx, id string) (model.Item, error) {
	return u.repo.GetItemById(c, id)
}

func (u usecase) UpdateItemById(c *fiber.Ctx, id, attribute string, value interface{}) error {
	return u.repo.UpdateItemById(c, id, attribute, value)
}

func (u usecase) DeleteItemById(c *fiber.Ctx, id string) error {
	return u.repo.DeleteItemById(c, id)
}

func (u usecase) PostItem(c *fiber.Ctx, item model.Item) error {
	return u.repo.PostItem(c, item)
}
