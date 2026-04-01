package mongodb

import (
	"context"
	"fmt"
	"go-mux-mongo-employees-manager/internal/domain/entities"
	"go-mux-mongo-employees-manager/internal/resources"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoRepo struct {
	Collection *mongo.Collection
}

func NewMongoRepo(collection *mongo.Collection) *MongoRepo {
	return &MongoRepo{
		Collection: collection,
	}
}

func UniqueIndexInMongoCollection(collection *mongo.Collection) error{
	model := mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), model)
	return err
}

func (repo *MongoRepo) Insert(model *entities.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.Collection.InsertOne(ctx, model)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%s", resources.EmailAlreadyExists)
		}
		return err
	}

	return nil
}
