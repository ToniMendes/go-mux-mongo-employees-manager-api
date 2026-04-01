// Package database provides database connection and initialization functionality.
package database

import "go.mongodb.org/mongo-driver/v2/mongo"

type database struct {
	MongoConn *mongo.Client
}

func NewDatabase(mongoConn *mongo.Client) *database {
	return &database{
		MongoConn: mongoConn,
	}
}
