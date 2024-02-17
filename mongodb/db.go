package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mscandan/url-shortener/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CreateClient(env *config.EnvironmentConfig) (*mongo.Client, error) {

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
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
