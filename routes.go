package main

import (
	"embajada-informer/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/item", handlers.GetItemDetails)
}
