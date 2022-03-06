package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChartRequest struct {
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	ProductId primitive.ObjectID `bson:"product_id" json:"product_id"`
}
