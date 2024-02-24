package service

import (
	"context"
	"log"
	"time"

	"github.com/mscandan/url-shortener/dto"
	"github.com/mscandan/url-shortener/model"
	"github.com/mscandan/url-shortener/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetFullUrlByShortUrl(db *mongo.Database, shortened_url string) (*model.UrlDocument, error) {
	var result model.UrlDocument
	err := db.Collection("urls").FindOne(context.TODO(), bson.M{"shortened_url": shortened_url}).Decode(&result)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result, err
}

func GetById(db *mongo.Database, id string) (*model.UrlDocument, error) {
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var doc model.UrlDocument
	err = db.Collection("urls").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&doc)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &doc, nil
}

func CreateShortUrl(db *mongo.Database, payload *dto.CreateShortenedUrlRequestPayload) (*string, error) {
	shortened_url := utils.GenerateRandomString(5)

	doc_to_insert := model.UrlDocument{FullUrl: payload.FullUrl, ShortenedUrl: shortened_url, CreatedAt: time.Now()}

	insert_result, err := db.Collection("urls").InsertOne(context.TODO(), &doc_to_insert)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	hex_id := insert_result.InsertedID.(primitive.ObjectID).Hex()

	return &hex_id, nil
}
