package database

import (
	"context"
	"fmt"
	"http_counter_service_api/internal/model"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
	GetIngredientById(id string) (model.Ingredient, error)
	GetIngredientByMenuId(id string) ([]model.Ingredient, error)
	GetIngredients() ([]model.Ingredient, error)
}

type service struct {
	db *mongo.Client
}

var (
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	//database = os.Getenv("DB_DATABASE")
)

func New() Service {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)))

	if err != nil {
		log.Fatal(err)
	}
	return &service{
		db: client,
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) GetIngredientById(id string) (model.Ingredient, error) {
	var ingredient model.Ingredient
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error converting id to object id. Err: %v", err)
	}

	result, err := s.db.Database("test").Collection("ingredients").Find(ctx, bson.M{"id": _id})
	if err != nil {
		fmt.Printf("error getting ingredient by id. Err: %v", err)
	}

	err = result.Decode(&ingredient)
	if err != nil {
		fmt.Printf("error decoding ingredient. Err: %v", err)
	}

	return ingredient, nil
}

func (s *service) GetIngredients() ([]model.Ingredient, error) {
	var ingredients []model.Ingredient
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cursor, err := s.db.Database("test").Collection("ingredients").Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var ingredient model.Ingredient
		err = cursor.Decode(&ingredient)
		if err != nil {
			fmt.Printf("error decoding ingredient. Err: %v", err)
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (s *service) GetIngredientByMenuId(id string) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error converting id to object id. Err: %v", err)
	}

	cursor, err := s.db.Database("test").Collection("ingredients").Find(ctx, bson.M{"menu_id": _id})

	for cursor.Next(ctx) {
		var ingredient model.Ingredient
		err = cursor.Decode(&ingredient)
		if err != nil {
			fmt.Printf("error decoding ingredient. Err: %v", err)
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}
