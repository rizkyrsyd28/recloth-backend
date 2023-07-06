package model

type Cart struct {
	UserId string `json:"id,omitempty" bson:"_id,omitempty"`
	List   []Item `json:"list" bson:"list"`
}
