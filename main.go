package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

const devEnvPort = "8080"

func main() {
	var port = isDev()

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.LoadHTMLGlob("./templates/*")
	SetupRoutes(r)

	if err := r.Run("0.0.0.0:" + port); err != nil {
		fmt.Printf("Error starting the server: %v", err)
		return
	}
}

// Returns PORT from environment if found, defaults to
// value in `port` parameter otherwise. The returned port
// is prefixed with a `:`, e.g. `":8080"`.
func isDev() string {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Main Error loading environmental variables: %s", err)
		return ""
	}
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return devEnvPort
}
