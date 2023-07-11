package model

type Item struct {
	Id             string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title          string   `json:"title" bson:"title"`
	Description    string   `json:"description" bson:"description"`
	Size           string   `json:"size" bson:"size"`
	Quantity       int      `json:"quantity" bson:"quantity"`
	Price          int      `json:"price" bson:"price"`
	ImgURL         []string `json:"img_url" bson:"img_url"`
	DonateDiscount float64  `json:"donate_discount" bson:"donate_discount"`
	Brand          string   `json:"brand" bson:"brand"`
	Location       string   `json:"location" bson:"location"`
	Condition      string   `json:"condition" bson:"condition"`
	OwnerId        string   `json:"owner_id,omitempty" bson:"owner_id,omitempty"`
}
