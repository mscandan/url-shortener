package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mscandan/url-shortener/dto"
	"github.com/mscandan/url-shortener/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (base *Controller) GetFullUrlByShortUrl(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.SendStatus(400)
	}

	// get from db if exists redirect to it
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}

	result, err := service.GetFullUrlByShortUrl(base.DB, objectId)

	if err != nil {
		log.Println(err)
		return err
	}

	return c.Redirect(result.FullUrl)
}

func (base *Controller) CreateShortUrl(c *fiber.Ctx) error {
	payload := new(dto.CreateShortenedUrlRequestPayload)

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	result, err := service.CreateShortUrl(base.DB, payload)

	if err != nil {
		log.Println(err)
		return err
	}

	formatted_short_url := c.Hostname() + "/" + *result

	span := "<span>" + formatted_short_url + "</span>"

	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	return c.Status(fiber.StatusCreated).SendString(span)
}
