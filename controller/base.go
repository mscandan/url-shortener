package controller

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	DB    *mongo.Database
	Cache *redis.Client
}
