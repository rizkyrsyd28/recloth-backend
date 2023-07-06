package model

type Item struct {
	Id             string  `json:"id,omitempty" bson:"_id,omitempty"`
	Title          string  `json:"title" bson:"title"`
	Description    string  `json:"description" bson:"description"`
	Size           string  `json:"size" bson:"size"`
	Quantity       int     `json:"quantity" bson:"quantity"`
	Price          int     `json:"price" bson:"price"`
	ImgURL         string  `json:"img_url" bson:"img_url"`
	DonateDiscount float32 `json:"donate_discount" bson:"donate_discount"`
}
