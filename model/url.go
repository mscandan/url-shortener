package model

import "time"

type UrlDocument struct {
	ShortenedUrl string    `bson:"shortened_url"`
	FullUrl      string    `bson:"full_url"`
	CreatedAt    time.Time `bson:"created_at"`
}
