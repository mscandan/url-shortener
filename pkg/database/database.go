package database

import (
	"context"
	"fmt"
	"log"

	"github.com/mscandan/url-shortener/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DB *mongo.Database
)

func createClient(env *config.EnvironmentConfig) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(env.MongoDB_URI))
	if err != nil {
		log.Println("Can't connect to Mongo", err)
		return nil, err
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		fmt.Printf("could not ping to mongo db service: %v\n", err)
		return nil, err
	}

	return client, nil
}

func Setup(env *config.EnvironmentConfig) error {
	client, err := createClient(env)

	if err != nil {
		return err
	}

	DB = client.Database(env.Database_Name)

	return nil
}

func GetDB() *mongo.Database {
	return DB
}
