package config

import (
	"context"
	"crud/env"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(env.GetValue("MONGO_URL")))
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
