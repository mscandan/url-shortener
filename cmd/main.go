package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mscandan/url-shortener/config"
	"github.com/mscandan/url-shortener/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	environment_config, err := config.Config()

	if err != nil {
		log.Fatalln(err)
	}

	mongoClient, err := mongodb.CreateClient(environment_config)

	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := mongoClient.Database(environment_config.Database_Name)
	collection := db.Collection("urls")

	res, err := collection.InsertOne(context.TODO(), bson.D{{Key: "pure_url", Value: "https://lmaoooo.com/lmao"}, {Key: "shortened_url", Value: "shortened.broski/123"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)

	if err != nil {
		log.Fatalln(err)
	}
}
