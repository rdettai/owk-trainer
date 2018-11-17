package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdettai/test-owkin/server/endpoints"
)

func generalHealthcheckHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func initRouter() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/v1/health", generalHealthcheckHandler())
	router.GET("/v1/models", endpoints.ListModels())
	router.POST("/v1/models", endpoints.SubmitModel())
	router.Run()
}

func main() {
	initRouter()
}
