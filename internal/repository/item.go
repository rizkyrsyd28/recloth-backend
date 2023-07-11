package repository

import (
	"errors"
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
	DecreaseItemQuantity(c *fiber.Ctx, itemId string, count int) error
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

	if err := cursor.All(c.Context(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r repo) GetItemById(c *fiber.Ctx, id string) (item model.Item, err error) {

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return item, err
	}

	err = r.DB.Collection("items").FindOne(c.Context(),
		&bson.D{{
			Key: "_id", Value: _id,
		}}).Decode(&item)

	return item, err
}

// UpdateItemById case sensitive, harus pas kalo mau update
func (r repo) UpdateItemById(c *fiber.Ctx, id, attribute string, value interface{}) (err error) {

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	opt := options.Update().SetUpsert(false)
	_, err = r.DB.Collection("items").UpdateOne(c.Context(), bson.M{"_id": _id}, bson.M{"$set": bson.M{attribute: value}}, opt)

	return err
}

func (r repo) DeleteItemById(c *fiber.Ctx, id string) (err error) {

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	del, err := r.DB.Collection("items").DeleteOne(c.Context(),
		bson.D{{
			"_id", _id,
		}})
	if err != nil {
		return err
	}

	if del.DeletedCount < 1 {
		err = errors.New("items not found")
	}

	return err
}

func (r repo) PostItem(c *fiber.Ctx, item model.Item) (err error) {

	_, err = r.DB.Collection("items").InsertOne(c.Context(), item)

	return err
}

func (r repo) DecreaseItemQuantity(c *fiber.Ctx, itemId string, count int) (err error) {

	_id, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}

	_, err = r.DB.Collection("items").UpdateOne(c.Context(),
		bson.M{"_id": _id},
		bson.M{"$inc": bson.M{"quantity": -count}})

	return err
}
