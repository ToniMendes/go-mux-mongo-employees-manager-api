package main

import (
	"context"
	"go-mux-mongo-employees-manager/internal/configs"
	"go-mux-mongo-employees-manager/internal/infra/database"
	"go-mux-mongo-employees-manager/internal/infra/database/mongodb"
	"go-mux-mongo-employees-manager/internal/usecase"
	"go-mux-mongo-employees-manager/internal/web"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	err := configs.StartConfig()
	returnFatalError(err)

	conn,coll := StartMongoDB()
	defer func() {
		err := conn.Disconnect(context.Background())
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	err = mongodb.UniqueIndexInMongoCollection(coll)
	returnFatalError(err)
	
	mongoRepo := mongodb.NewMongoRepo(coll)

	createUseCase := usecase.NewCreateUseCase(mongoRepo)

	type handler struct {
		*usecase.CreateUseCase
	}

	hub := handler{
		createUseCase,
	}

	usecase := web.NewHandler(hub)

	err = web.Routers(usecase, configs.Env.PORT)
	returnFatalError(err)
}

func StartMongoDB() (*mongo.Client, *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := mongodb.NewMongoConn(ctx, configs.Env.MongoURI)
	returnFatalError(err)

	client := database.NewDatabase(conn)

	return conn, mongodb.GetCollection(client.MongoConn, configs.Env.MongoDB, configs.Env.MongoCollection)
}

func returnFatalError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
