package main

import (
	"WordbookGenerater-Go/backend/api"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/", "./frontend/static")
	apiGroup := r.Group("/api")

	api.RegisterWordbook(apiGroup)
	api.RegisterWordTest(apiGroup)

	log.Println("Server started at :8080")
	r.Run(":8080")
}
