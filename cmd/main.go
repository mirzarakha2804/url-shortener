package main

import (
	"log"
	"url-shortener/config"
	"url-shortener/internal/handler"
	"url-shortener/repository"

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
	repository.NewURLRepository(db)

	r := gin.Default()

	r.GET("/", handler.HelloHandler)

	r.Run(":8080")
}
