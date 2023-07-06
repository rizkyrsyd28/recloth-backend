package model

import "time"

type Transaction struct {
	Id       string    `json:"id,omitempty" bson:"_id,omitempty"`
	Date     time.Time `json:"date" bson:"date"`
	Quantity int       `json:"quantity" bson:"quantity"`
	Items    Item      `json:"items" bson:"items"`
	Amount   int       `json:"amount" bson:"amount"`
}
