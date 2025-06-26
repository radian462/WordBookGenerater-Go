package main

import (
	"log"
	"WordbookGenerater-Go/backend/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/", "./frontend/static")
	apiGroup := r.Group("/api")

	api.RegisterWordbook(apiGroup)

	log.Println("Server started at :8080")
	r.Run(":8080")
}
