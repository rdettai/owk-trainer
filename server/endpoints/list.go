package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rdettai/test-owkin/server/conf"
)

type Score struct {
	Model    string      `json:"name"`
	Loss     interface{} `json:"test_loss"`
	Accuracy interface{} `json:"test_accuracy"`
}

func loadScore(file string) Score {
	var score Score
	scoreFile, err := os.Open(file)
	defer scoreFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(scoreFile)
	jsonParser.Decode(&score)
	score.Model = file
	return score
}

func ListModels() func(c *gin.Context) {
	return func(c *gin.Context) {
		var files []Score

		err := filepath.Walk(conf.ModelFolder, func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".json") {
				files = append(files, loadScore(path))
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"scores": files,
		})
	}
}
