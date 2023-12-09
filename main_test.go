package main

import (
	"embajada-informer/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetItemDetails(t *testing.T) {
	// Set up a Gin router
	router := gin.Default()
	SetupRoutes(router)

	// Create a request to the endpoint
	req, err := http.NewRequest("GET", "/item", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Serve the request to the router
	router.ServeHTTP(w, req)

	// Your JSON response body
	responseBody := w.Body

	// Create an instance of your defined type
	var response model.Response

	// Unmarshal the JSON response into the defined type
	err = json.Unmarshal(responseBody.Bytes(), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Now 'response' contains the data from the JSON response
	fmt.Printf("'response' contains the data from the JSON response  %+v\n", responseBody)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the expected response
	expectedType := reflect.TypeOf(model.Response{})
	actualType := reflect.TypeOf(response)

	assert.Equal(t, expectedType, actualType)
}
