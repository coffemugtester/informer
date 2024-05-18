package main

import (
	"embajada-informer/internal/handlers"
	"embajada-informer/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/index", handlers.GetVIP)
	if port := isDev(); port == devEnvPort {
		r.GET("/mock", func(c *gin.Context) {
			fmt.Printf("Running mock endpoint\n")
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
