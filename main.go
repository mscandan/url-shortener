package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mscandan/url-shortener/controller"
	"github.com/mscandan/url-shortener/pkg/cache"
	"github.com/mscandan/url-shortener/pkg/config"
	"github.com/mscandan/url-shortener/pkg/database"
)

func main() {
	environment_config, err := config.Config()
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Setup(environment_config); err != nil {
		log.Fatalln(err)
	}

	if err := cache.Setup(environment_config); err != nil {
		log.Fatalln(err)
	}

	db := database.GetDB()
	cache_client := cache.GetClient()

	controller := controller.Controller{DB: db, Cache: cache_client}

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

	app.Get("/:shortened_url", controller.GetFullUrlByShortUrl)

	app.Post("/url", controller.CreateShortUrl)

	app.Listen(":" + environment_config.Port)

}
