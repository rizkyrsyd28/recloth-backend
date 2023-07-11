package model

import "time"

type Transaction struct {
	Id       string    `json:"id,omitempty" bson:"_id,omitempty"`
	Date     time.Time `json:"date" bson:"date"`
	Quantity int       `json:"quantity" bson:"quantity"`
	ItemId   string    `json:"item_id,omitempty" bson:"item_id,omitempty"`
	UserId   string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Amount   int       `json:"amount" bson:"amount"`
}
