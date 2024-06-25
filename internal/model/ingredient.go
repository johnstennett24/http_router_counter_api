package model

type Ingredient struct {
	_id         string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
}
