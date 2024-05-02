package main

import (
	"embajada-informer/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	if port := isDev(); port == "8080" {
		r.GET("/mock", func(c *gin.Context) {
			//TODO: return mock data
			c.JSON(http.StatusOK, "hello")
		})
	} else {
		r.GET("/item", handlers.GetItemDetails)
	}
}
