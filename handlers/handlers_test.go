package handlers

import (
	"auth-system/database"
	"auth-system/models"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() {
	// Initialize the database
	database.InitDB()

	// Auto-migrate the User model
	database.DB.AutoMigrate(&models.User{})
}

func teardown() {
	// Drop the User table after the test
	database.DB.Migrator().DropTable(&models.User{})
}

func TestSignup(t *testing.T) {
	// Setup the database
	setup()
	defer teardown() // Clean up after the test

	// Create a request body
	requestBody := bytes.NewBufferString(`{"email": "test@example.com", "password": "password123"}`)

	// Create a request
	req, err := http.NewRequest("POST", "/signup", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler directly
	handler := http.HandlerFunc(Signup)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
	}

	// Check the response body
	expected := `{"message":"User created successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}