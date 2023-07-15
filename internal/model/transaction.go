package model

import "time"

type Transaction struct {
	Id     string     `json:"id,omitempty" bson:"_id,omitempty"`
	Date   time.Time  `json:"date" bson:"date"`
	Items  []CartItem `json:"items,omitempty" bson:"items,omitempty"`
	UserId string     `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Amount int        `json:"amount" bson:"amount"`
}
