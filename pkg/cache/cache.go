package cache

import (
	"log"

	"github.com/mscandan/url-shortener/pkg/config"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func Setup(env *config.EnvironmentConfig) error {
	log.Println("Attempting to create redis client")
	opt, err := redis.ParseURL(env.Redis_URL)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	RedisClient = redis.NewClient(opt)

	log.Println("redis client creation succesfull")

	return nil
}

func GetClient() *redis.Client {
	return RedisClient
}
