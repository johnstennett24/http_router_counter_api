package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	id      primitive.ObjectID `bson:"_id"`
	storeid string             `bson:"storeid"`
	name    string             `bson:"name"`
	menu    Menu
}
