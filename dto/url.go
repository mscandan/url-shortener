package dto

type CreateShortenedUrlRequestPayload struct {
	FullUrl string `json:"full_url"`
}

type CreateShortenedUrlResponse struct {
	ShortLink string `json:"short_link"`
}
