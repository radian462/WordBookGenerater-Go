package api

import (
	"WordbookGenerater-Go/backend/pkg"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
)

type WordTestQuery struct {
	BaseWordbookPath string `json:"baseWordbookPath" binding:"required"`
	Rng              string `json:"rng" binding:"required"`
	IsReverse        bool   `json:"isReverse"`
	IsRandom         bool   `json:"isRandom"`
}

func RegisterWordTest(r *gin.RouterGroup) {
	r.POST("/wordtest", func(c *gin.Context) {
		var query WordTestQuery

		if err := c.ShouldBindJSON(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
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

		file, err := os.Open(query.BaseWordbookPath)
		if err != nil {
			log.Printf("File open failed: path=%s, error=%v", query.BaseWordbookPath, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":   500,
				"filepath": "",
			})
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		rows, err := reader.ReadAll()
		if err != nil {
			log.Printf("CSV read failed: path=%s, error=%v", query.BaseWordbookPath, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":   500,
				"filepath": "",
			})
			return
		}

		indices := make([]int, len(rng))
		copy(indices, rng)

		if query.IsRandom {
			rand.Shuffle(len(indices), func(i, j int) {
				indices[i], indices[j] = indices[j], indices[i]
			})
		}

		var words, answers []string
		for _, i := range indices {
			if i-1 < 0 || i-1 >= len(rows) {
				log.Printf("Index out of bounds: i=%d, rows=%d", i, len(rows))
				continue
			}
			if len(rows[i-1]) < 2 {
				log.Printf("Invalid row format at line %d: %#v", i, rows[i-1])
				continue
			}
			words = append(words, rows[i-1][0])
			answers = append(answers, rows[i-1][1])
		}

		if query.IsReverse {
			words, answers = answers, words
		}

		base := filepath.Base(query.BaseWordbookPath)
		ext := filepath.Ext(base)
		outName := pkg.NameGenerate(base[:len(base)-len(ext)], "xlsx", rng, 207, "resources/output")
		path, err := pkg.GenerateWordTest(
			"resources/template.xlsx",
			"resources/output",
			outName,
			words,
			answers,
		)
		if err != nil {
			log.Printf("GenerateWordTest failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":   500,
				"filepath": "",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   200,
			"filepath": "/" + path,
		})
	})
}
