package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
)

type TransactionUsecase interface {
	GetTransactionByUserId(c *fiber.Ctx, userId string) ([]model.Transaction, error)
	Checkout(c *fiber.Ctx, transaction model.Transaction) error
}

func (u usecase) GetTransactionByUserId(c *fiber.Ctx, userId string) ([]model.Transaction, error) {
	return u.repo.GetTransactionByUserId(c, userId)
}

func (u usecase) Checkout(c *fiber.Ctx, transaction model.Transaction) (err error) {

	err = u.repo.DecreaseItemQuantity(c, transaction.ItemId, transaction.Quantity)
	if err != nil {
		return err
	}

	if err = u.repo.DecreaseBalance(c, transaction.UserId, transaction.Amount); err != nil {
		return err
	}

	err = u.repo.PostTransaction(c, transaction)

	return err
}
