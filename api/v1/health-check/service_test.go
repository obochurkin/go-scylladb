package healthCheck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/obochurkin/go-scylladb/api/v1/health-check" // Update this import path
)

func TestHealthCheckHandler(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a response recorder to record the response
	rec := httptest.NewRecorder()

	// Create a new Gin context using the recorder and request
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Call the handler function
	healthCheck.HealthCheckHandler(c)

	// Check the response status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expectedBody := `{"status":200}`
	if rec.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expectedBody)
	}
}
