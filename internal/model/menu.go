package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Menu struct {
	id    primitive.ObjectID `bson:"_id"`
	items []MenuItem         `bosn:"items"`
}
