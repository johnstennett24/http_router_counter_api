package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MenuItem struct {
	id          primitive.ObjectID `bson:"_id"`
	name        string             `bson:"name"`
	price       float64            `bson:"price"`
	cost        float32            `bson:"cost"`
	ingredients []Ingredient       `bson:"ingredients"`
}
