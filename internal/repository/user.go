package repository

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepo interface {
	GetUserByUsername(c *fiber.Ctx, username string) (model.User, error)
	CreateUser(c *fiber.Ctx, user model.User) (string, error)
	DeleteUser(c *fiber.Ctx, username string) error
	DecreaseBalance(c *fiber.Ctx, userId string, amount int) error
}

func (r repo) GetUserByUsername(c *fiber.Ctx, username string) (user model.User, err error) {

	err = r.DB.Collection("users").FindOne(c.Context(),
		bson.D{{
			Key: "username", Value: username,
		}}).Decode(&user)

	return user, err
}

func (r repo) CreateUser(c *fiber.Ctx, user model.User) (id string, err error) {

	ins, err := r.DB.Collection("users").InsertOne(c.Context(), user)
	if err != nil {
		return id, err
	}

	id = ins.InsertedID.(primitive.ObjectID).Hex()

	return id, err
}

func (r repo) DeleteUser(c *fiber.Ctx, username string) error {

	del, err := r.DB.Collection("users").DeleteOne(c.Context(),
		bson.D{{
			"username", username,
		}})
	if err != nil {
		return err
	}
	if del.DeletedCount < 1 {
		return errors.New("internal server error")
	}

	return nil
}

func (r repo) DecreaseBalance(c *fiber.Ctx, userId string, amount int) error {

	_id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	var user model.User

	err = r.DB.Collection("users").FindOne(c.Context(),
		bson.D{{
			Key: "_id", Value: _id,
		}}).Decode(&user)
	if err != nil {
		return err
	}

	if amount > user.Balance {
		return errors.New("saldo tidak cukup")
	} else {
		result, err := r.DB.Collection("users").UpdateOne(c.Context(),
			bson.M{"_id": _id},
			bson.M{"$inc": bson.M{"balance": -amount}})
		if err != nil {
			return err
		}

		if result.ModifiedCount < 1 {
			return errors.New("cant modify")
		}
	}

	return err
}
