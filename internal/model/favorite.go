package model

type Favorite struct {
	UserId string `json:"id,omitempty" bson:"_id,omitempty"`
	List   []Item `json:"list" bson:"list"`
}
