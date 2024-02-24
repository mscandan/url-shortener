package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mscandan/url-shortener/dto"
	"github.com/mscandan/url-shortener/service"
)

func (base *Controller) GetFullUrlByShortUrl(c *fiber.Ctx) error {
	shortened_url := c.Params("shortened_url")

	if shortened_url == "" {
		return c.SendStatus(400)
	}

	result, err := service.GetFullUrlByShortUrl(base.DB, base.Cache, shortened_url)

	if err != nil {
		log.Println(err)
		return err
	}

	return c.Redirect(*result)
}

func (base *Controller) CreateShortUrl(c *fiber.Ctx) error {
	payload := new(dto.CreateShortenedUrlRequestPayload)

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	created_id, err := service.CreateShortUrl(base.DB, base.Cache, payload)

	if err != nil {
		log.Println(err)
		return err
	}

	created_doc, err := service.GetById(base.DB, *created_id)

	if err != nil {
		log.Println(err)
		return err
	}

	formatted_short_url := c.Hostname() + "/" + created_doc.ShortenedUrl

	span := "<span>" + formatted_short_url + "</span>"

	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	return c.Status(fiber.StatusCreated).SendString(span)
}
