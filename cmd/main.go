package main

import (
	"log"
	"os"
	"url-shortener/config"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env variables")
	}

	// Connect to database
	db := config.ConnectDB()
	defer db.Close()
	log.Println("✅ Database connection established")
	// Setup repository
	shortRepo := repository.NewURLRepository(db)
	shortService := service.NewURLRepository(shortRepo)
	shortHandler := handler.NewUrlHandler(shortService)
	r := gin.Default()

	r.GET("/", handler.HelloHandler)
	r.POST("/shorten", shortHandler.ShortenURL)

	r.Run(os.Getenv("APP_PORT"))
}
