package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
)

type CartUsecase interface {
	CreateCart(c *fiber.Ctx, userId string) error
	GetCart(c *fiber.Ctx, userId string) ([]model.Item, error)
	UpdateCart(c *fiber.Ctx, method, itemId, userId string) error
	ClearCart(c *fiber.Ctx, userId string) error
}

func (u usecase) CreateCart(c *fiber.Ctx, userId string) (err error) {
	return u.repo.CreateCart(c, userId)
}

func (u usecase) GetCart(c *fiber.Ctx, userId string) ([]model.Item, error) {

	cart, err := u.repo.GetCartByUserId(c, userId)
	if err != nil {
		return nil, err
	}

	result := make([]model.Item, 0)

	for _, cartItem := range cart.List {

		item, err := u.repo.GetItemById(c, cartItem.ItemId)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}

func (u usecase) UpdateCart(c *fiber.Ctx, method, itemId, userId string) error {

	fmt.Println("PASS1")
	cart, err := u.repo.GetCartByUserId(c, userId)
	if err != nil {
		return err
	}

	fmt.Println("PASS2")
	if err = u.repo.UpdateCart(c, method, itemId, cart.Id); err != nil {
		return err
	}

	return nil
}

func (u usecase) ClearCart(c *fiber.Ctx, userId string) error {

	cart, err := u.repo.GetCartByUserId(c, userId)
	if err != nil {
		return err
	}

	if err := u.repo.ClearCart(c, cart.Id); err != nil {
		return err
	}

	return nil
}
