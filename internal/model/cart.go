package model

type Cart struct {
	Id     string     `json:"id,omitempty" bson:"_id,omitempty"`
	UserId string     `json:"user_id,omitempty" bson:"user_id,omitempty"`
	List   []CartItem `json:"list,omitempty" bson:"list,omitempty"`
	//List   []string `json:"list,omitempty" bson:"list,omitempty"`
}

type CartItem struct {
	Quantity int    `json:"quantity" bson:"quantity"`
	ItemId   string `json:"item_id,omitempty" bson:"item_id,omitempty"`
}
