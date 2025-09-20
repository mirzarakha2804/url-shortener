package main

import (
	"url-shortener/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handler.HelloHandler)

	r.Run(":8080")
}
