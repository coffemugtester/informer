package handlers

import (
	"embajada-informer/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetItemDetails(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading environmental variables")
		return
	}
	apiKey := os.Getenv("API_KEY")
	externalAPIURL := "https://maps.googleapis.com/maps/api/place/details/json?fields=name%2Crating%2Copening_hours%2Cwebsite%2Caddress_components%2Cadr_address%2Cbusiness_status%2Cformatted_address%2Cformatted_phone_number%2Cgeometry%2Crating%2Cuser_ratings_total%2Creviews%2Copening_hours%2Cphotos&reviews_sort=newest&reviews_no_translations=true&place_id=ChIJs6C7pbRRqEcRu0jebrF6XuQ&key=" + apiKey

	// Make a GET request to the external API
	response, err := http.Get(externalAPIURL)
	if err != nil {
		// Handle the error and respond with an appropriate error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error making external API request: %v", err)})
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		// Handle the error and respond with an appropriate error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create an instance of your struct to unmarshal into
	var item model.Response

	// Unmarshal the JSON response into the struct
	if err := json.Unmarshal(body, &item); err != nil {
		// Handle the decoding error and respond with an appropriate error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("THE UNMARSHALLED BODY: %v", item)

	// Respond with the decoded item or perform further processing
	c.JSON(http.StatusOK, item)
}
