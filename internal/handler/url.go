package handler

import (
	"net/http"
	"url-shortener/internal/dto"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service *service.ShortenerService
}

func NewUrlHandler(service *service.ShortenerService) *handler {
	return &handler{
		service: service,
	}
}

func HelloHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello")
}

func (h *handler) ShortenURL(ctx *gin.Context) {
	var requestBody struct {
		URL string `json:"url" binding:"required,min=1"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	body := dto.ShortenURLRequest{
		URL: requestBody.URL,
	}
	shortUrl, err := h.service.ShortenURL(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &shortUrl)
}
