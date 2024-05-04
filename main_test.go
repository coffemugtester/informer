package main

import (
	"embajada-informer/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetItemsDetails(t *testing.T) {
	t.Run("item handler returns html", func(t *testing.T) {

		gin.SetMode(gin.TestMode)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/item", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		r := gin.Default()

		r.GET("/item", handlers.GetItemDetails)

		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

	})
}
