package main

import (
	"embajada-informer/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	if port := envPortOr("8080"); port == "8080" {
		r.GET("/item", handlers.GetItemDetails)
	} else {
		r.GET("/mock", func(c *gin.Context) {
			//TODO: return mock data
			c.JSON(http.StatusOK, "hello")
		})
	}
}
