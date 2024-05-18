package main

import (
	"embajada-informer/internal/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVIP(t *testing.T) {
	t.Run("VIP handler returns HTML", func(t *testing.T) {

		r := gin.Default()

		r.LoadHTMLGlob("templates/*")

		rr := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/vip", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		r.GET("/vip", handlers.GetVIP)

		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status OK; got %v", rr.Code)
		}
		if rr.Header().Get("Content-Type") != "text/html; charset=utf-8" {
			t.Fatalf("unexpected content type: %v", rr.Header().Get("Content-Type"))
		}
		fmt.Printf("Response Body:\n%v\n", rr.Body.String())
	})
}

func TestGetItemsDetails(t *testing.T) {
	t.Run("item handler returns StatusOk", func(t *testing.T) {

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
