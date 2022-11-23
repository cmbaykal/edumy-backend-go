package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "EdumyDB"
const mongoURI = "mongodb://localhost:27017"

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	fmt.Println("Successfully connected to the database.")
	return nil
}

func GetInstance() MongoInstance {
	return mg
}


// TODO: add getCollection method