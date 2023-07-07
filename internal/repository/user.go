package repository

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepo interface {
	GetUserByUsername(c *fiber.Ctx, username string) (model.User, error)
	CreateUser(c *fiber.Ctx, user model.User) error
	DeleteUser(c *fiber.Ctx, username string) error
}

func (r repo) GetUserByUsername(c *fiber.Ctx, username string) (user model.User, err error) {

	err = r.DB.Collection("users").FindOne(c.Context(),
		bson.D{{
			Key: "username", Value: username,
		}}).Decode(&user)

	return user, err
}

func (r repo) CreateUser(c *fiber.Ctx, user model.User) (err error) {

	_, err = r.DB.Collection("users").InsertOne(c.Context(), user)

	return err
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
