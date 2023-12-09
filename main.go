package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

// Returns PORT from environment if found, defaults to
// value in `port` parameter otherwise. The returned port
// is prefixed with a `:`, e.g. `":3000"`.
func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Main Error loading environmental variables")
		return
	}

	// ...
	// Use `PORT` provided in environment or default to 8080
	var port = envPortOr("8080")

	r := gin.Default()
	SetupRoutes(r)

	// Listen on all available network interfaces (0.0.0.0) on the specified port
	if err := r.Run("0.0.0.0" + port); err != nil {
		fmt.Printf("Error starting the server: %v", err)
		return
	}

}
