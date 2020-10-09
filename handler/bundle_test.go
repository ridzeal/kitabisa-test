package handler

import (
	"bytes"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"

	"kitabisa-test/entity"
)

func TestBundlePositive(t *testing.T) {
	var jsonStr = []byte(`{"cakes":10,"apples":10}`)

	req, err := http.NewRequest("POST", "/bundle", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Bundle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestBundleNotJson(t *testing.T) {
	var jsonStr = []byte(`{"cakes":10,"apples":10`)

	req, err := http.NewRequest("POST", "/bundle", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Bundle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestBundleMissingParam(t *testing.T) {
	var jsonStr = []byte(`{"cakes":10}`)

	req, err := http.NewRequest("POST", "/bundle", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Bundle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	var response *entity.BundleResponse
	json.Unmarshal(rr.Body.Bytes(), &response)
	expectedMessage := "Zero quantity"
	if response.ErrorMessage != expectedMessage {
		t.Errorf("handler returned unexpected error message: got %v want %v",
			response.ErrorMessage, expectedMessage)
	}
}