package main

import (
	"embajada-informer/internal/handlers"
	"embajada-informer/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	if port := isDev(); port == "8080" {
		r.GET("/mock", func(c *gin.Context) {
			response := model.Response{
				HTMLAttributions: nil,
				Result:           model.Result{},
			}
			c.JSON(http.StatusOK, response)
		})
	} else {
		r.GET("/item", handlers.GetItemDetails)
	}
}
