package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartRepo interface {
	CreateCart(c *fiber.Ctx, userId string) error
	GetCartByUserId(c *fiber.Ctx, userId string) (model.Cart, error)
	UpdateCart(c *fiber.Ctx, method, itemId, cartId string) error
	ClearCart(c *fiber.Ctx, id string) error
}

func (r repo) CreateCart(c *fiber.Ctx, userId string) (err error) {

	cart := model.Cart{
		Id:     "",
		UserId: userId,
		List:   make([]model.CartItem, 0),
	}

	_, err = r.DB.Collection("carts").InsertOne(c.Context(), cart)

	return err
}

func (r repo) GetCartByUserId(c *fiber.Ctx, userId string) (cart model.Cart, err error) {

	err = r.DB.Collection("carts").FindOne(c.Context(),
		&bson.D{{
			Key: "user_id", Value: userId,
		}}).Decode(&cart)

	return cart, err
}

func (r repo) UpdateCart(c *fiber.Ctx, method, itemId, cartId string) error {

	fmt.Println("CART1")
	_id, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		return err
	}

	if method == "add" {
		method = "$push"
	} else {
		if method == "del" {
			method = "$pull"
		} else {
			return errors.New("invalid method")
		}
	}

	_, err = r.DB.Collection("carts").UpdateOne(c.Context(),
		bson.M{"_id": _id},
		bson.M{method: bson.M{"list": model.CartItem{ItemId: itemId, Quantity: 1}}},
	)
	if err != nil {
		return err
	}

	fmt.Println("CART2")

	return nil
}

func (r repo) ClearCart(c *fiber.Ctx, cartId string) error {

	_id, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		return err
	}

	_, err = r.DB.Collection("carts").UpdateOne(c.Context(),
		bson.M{
			"_id": _id,
		},
		bson.M{
			"$set": bson.M{
				"list": make([]model.CartItem, 0),
			},
		})
	if err != nil {
		return err
	}

	return nil
}
