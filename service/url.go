package service

import (
	"context"
	"log"
	"time"

	"github.com/mscandan/url-shortener/dto"
	"github.com/mscandan/url-shortener/model"
	"github.com/mscandan/url-shortener/utils"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// make this return only string
func GetFullUrlByShortUrl(db *mongo.Database, cache *redis.Client, shortened_url string) (*string, error) {
	// try to get from  cache
	cached_value, err := cache.Get(context.Background(), shortened_url).Result()

	if err != nil {
		log.Println("value not found for key: ", shortened_url, " in cache")

		var result model.UrlDocument
		err = db.Collection("urls").FindOne(context.TODO(), bson.M{"shortened_url": shortened_url}).Decode(&result)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &result.FullUrl, err
	}

	return &cached_value, nil

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

func CreateShortUrl(db *mongo.Database, cache *redis.Client, payload *dto.CreateShortenedUrlRequestPayload) (*string, error) {
	shortened_url := utils.GenerateRandomString(5)

	doc_to_insert := model.UrlDocument{FullUrl: payload.FullUrl, ShortenedUrl: shortened_url, CreatedAt: time.Now()}

	insert_result, err := db.Collection("urls").InsertOne(context.TODO(), &doc_to_insert)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = cache.Set(context.Background(), shortened_url, payload.FullUrl, 10*time.Minute).Err()

	if err != nil {
		log.Println("failed to write to cache", err)
	}

	hex_id := insert_result.InsertedID.(primitive.ObjectID).Hex()

	return &hex_id, nil
}
