package dto

type ShortenURLRequest struct {
	URL string `json:"url" binding:"required,min=1"`
}
