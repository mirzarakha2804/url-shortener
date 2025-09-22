package service

import (
	"fmt" // Update this path if your dto is located elsewhere, e.g. "url-shortener/internal/dto/request"
	"url-shortener/internal/dto"
	"url-shortener/internal/repository"
	"url-shortener/pkg/utils"
)

type ShortenerService struct {
	repo *repository.URLRepository
}

func NewURLRepository(repository *repository.URLRepository) *ShortenerService {
	return &ShortenerService{
		repo: repository,
	}
}

func (r *ShortenerService) ShortenURL(originalUrl dto.ShortenURLRequest) (*dto.ShortenURLResponse, error) {
	if r.repo.IsUrlExist(originalUrl.URL) {
		return nil, fmt.Errorf("url already exists")
	}
	idGet, err := r.repo.GetNextId()
	if err != nil {
		return nil, fmt.Errorf("failed to get next id: %w", err)
	}
	shortUrl := utils.Encode(idGet)

	//add to db AddUrl

	err = r.repo.AddUrl(originalUrl.URL, shortUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create url: %w", err)
	}

	response := &dto.ShortenURLResponse{
		ShortURL: shortUrl,
	}
	return response, nil
}
