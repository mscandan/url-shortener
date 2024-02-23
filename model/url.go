package model

import "time"

type UrlDocument struct {
	FullUrl   string    `bson:"full_url"`
	CreatedAt time.Time `bson:"created_at"`
}
