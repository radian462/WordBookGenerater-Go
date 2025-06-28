package main

import (
	"WordbookGenerater-Go/backend/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.MkdirAll("resources/output", 0755)
	r := gin.Default()

	r.Static("/static", "./frontend/static")
	r.Static("resources/output", "./resources/output")
	apiGroup := r.Group("/api")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(308, "/static/index.html")
	})

	api.RegisterWordbook(apiGroup)
	api.RegisterWordTest(apiGroup)

	log.Println("Server started at :8080")
	r.Run("0.0.0.0:8080")
}
