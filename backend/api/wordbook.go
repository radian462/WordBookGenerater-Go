package api

import (
	"WordbookGenerater-Go/backend/pkg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type WordbookQuery struct {
	BaseWordbookPath string `json:"baseWordbookPath" binding:"required"`
	Rng              string `json:"rng" binding:"required"`
}

func RegisterWordbook(r *gin.RouterGroup) {
	r.POST("/wordbook", func(c *gin.Context) {
		var query WordbookQuery

		if err := c.ShouldBindJSON(&query); err != nil {
			c.JSON(http.StatusPaymentRequired, gin.H{
				"status":   400,
				"filepath": "",
			})
			return
		}

		rng, err := pkg.RangeParse(query.Rng)
		if err != nil {
			log.Printf("RangeParse failed: input=%s, error=%v", query.Rng, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":   500,
				"filepath": "",
			})
			return
		}

		path, err := pkg.ExtractCSVRange(
			query.BaseWordbookPath,
			"resources/output",
			"",
			rng,
			[]string{"FrontText", "BackText", "Comment", "FrontTextLanguage", "BackTextLanguage"},
		)

		if err != nil {
			log.Printf("ExtractCSVRange failed: error=%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":   500,
				"filepath": "",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   200,
			"filepath": path,
		})
	})
}
