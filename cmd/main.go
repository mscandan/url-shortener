package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mscandan/url-shortener/config"
	"github.com/mscandan/url-shortener/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateShortenedUrlRequestPayload struct {
	FullUrl string `json:"full_url"`
}

type CreateShortenedUrlResponse struct {
	ShortLink string `json:"short_link"`
}

type UrlDoc struct {
	FullUrl   string    `bson:"full_url"`
	CreatedAt time.Time `bson:"created_at"`
}

func main() {
	environment_config, err := config.Config()
	if err != nil {
		log.Fatalln(err)
	}

	mongoClient, err := mongodb.CreateClient(environment_config)
	if err != nil {
		log.Fatalln(err)
	}

	db := mongoClient.Database(environment_config.Database_Name)

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	app := fiber.New()

	// serve static files
	app.Static("/", "./static", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 60 * time.Minute,
		MaxAge:        3600,
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// get from db if exists redirect to it
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println(err)
			return err
		}

		var result UrlDoc
		err = db.Collection("urls").FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&result)

		if err != nil {
			log.Println(err)
			return err
		}

		return c.Redirect(result.FullUrl)
	})

	app.Post("/url", func(c *fiber.Ctx) error {
		payload := new(CreateShortenedUrlRequestPayload)

		if err := c.BodyParser(payload); err != nil {
			return err
		}

		doc := UrlDoc{FullUrl: payload.FullUrl, CreatedAt: time.Now()}

		// save these to db
		result, err := db.Collection("urls").InsertOne(context.TODO(), &doc)

		if err != nil {
			log.Println(err)
			return err
		}

		response := CreateShortenedUrlResponse{
			ShortLink: c.Hostname() + "/" + result.InsertedID.(primitive.ObjectID).Hex(),
		}

		return c.Status(fiber.StatusCreated).JSON(&response)
	})

	app.Listen(":8080")

}
