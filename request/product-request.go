package request

type ProductRequest struct {
	Tilte      string  `bson:"title" json:"title"`
	PictureUrl string  `bson:"picture_url" json:"picture_url"`
	Price      float32 `bson:"price" json:"price"`
}
