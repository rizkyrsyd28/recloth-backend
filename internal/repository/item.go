package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ItemRepo interface {
	GetAllItems(c *fiber.Ctx, page, limit int) ([]model.Item, error)
	GetItemById(c *fiber.Ctx, id string) (model.Item, error)
	UpdateItemById(c *fiber.Ctx, id, attribute string, value interface{}) error
	DeleteItemById(c *fiber.Ctx, id string) error
	PostItem(c *fiber.Ctx, item model.Item) error
}

func (r repo) GetAllItems(c *fiber.Ctx, page, limit int) ([]model.Item, error) {
	result := make([]model.Item, 0)

	opt := options.Find().SetLimit(int64(limit)).SetSkip(int64((page - 1) * limit))
	cursor, err := r.DB.Collection("items").Find(c.Context(),
		bson.M{"quantity": bson.M{"$gt": 0}}, opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c.Context())

	for cursor.Next(c.Context()) {
		var item model.Item
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, nil
}

func (r repo) GetItemById(c *fiber.Ctx, id string) (item model.Item, err error) {

	_id, err := primitive.ObjectIDFromHex(id)

	err = r.DB.Collection("items").FindOne(c.Context(),
		bson.D{{
			Key: "_id", Value: _id,
		}}).Decode(&item)

	return item, err
}
