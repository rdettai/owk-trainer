package endpoints

import (
	"net/http"

	"github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

func SubmitModel() func(c *gin.Context) {
	return func(c *gin.Context) {
		endpoint := "unix:///var/run/docker.sock"
		client, err := docker.NewClient(endpoint)
		if err != nil {
			panic(err)
		}
		imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"images": imgs,
		})
	}
}
