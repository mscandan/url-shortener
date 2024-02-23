package service

import (
	"context"
	"log"
	"time"

	"github.com/mscandan/url-shortener/dto"
	"github.com/mscandan/url-shortener/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetFullUrlByShortUrl(db *mongo.Database, objectId primitive.ObjectID) (*model.UrlDocument, error) {
	var result model.UrlDocument
	err := db.Collection("urls").FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&result)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result, err
}

func CreateShortUrl(db *mongo.Database, payload *dto.CreateShortenedUrlRequestPayload) (*string, error) {
	doc := model.UrlDocument{FullUrl: payload.FullUrl, CreatedAt: time.Now()}

	result, err := db.Collection("urls").InsertOne(context.TODO(), &doc)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()

	return &insertedId, err
}
