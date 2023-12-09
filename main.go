package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	SetupRoutes(r)
	r.Run(":8080")

	fmt.Printf("Server listing on port 8080...")
}
