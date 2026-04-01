package mongodb

import (
	"context"
	"go-mux-mongo-employees-manager/internal/domain/entities"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoRepo struct {
	Collection *mongo.Collection
}

func NewMongoRepo(collection *mongo.Collection) *MongoRepo {
	return &MongoRepo{
		Collection: collection,
	}
}

func (repo *MongoRepo) Insert(model *entities.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.Collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
