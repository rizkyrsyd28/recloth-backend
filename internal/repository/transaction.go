package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionRepo interface {
	GetTransactionByUserId(c *fiber.Ctx, userId string) ([]model.Transaction, error)
	PostTransaction(c *fiber.Ctx, transaction model.Transaction) error
}

func (r repo) GetTransactionByUserId(c *fiber.Ctx, userId string) ([]model.Transaction, error) {

	result := make([]model.Transaction, 0)

	opt := options.Find().SetSort(bson.D{{"date", -1}})
	cursor, err := r.DB.Collection("transactions").Find(c.Context(),
		bson.M{"user_id": userId}, opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c.Context())

	if err := cursor.All(c.Context(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r repo) PostTransaction(c *fiber.Ctx, transaction model.Transaction) (err error) {

	_, err = r.DB.Collection("transactions").InsertOne(c.Context(), transaction)

	return err
}
