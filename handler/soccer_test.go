package handler

import (
	"bytes"
	"testing"
	"net/http"
	"net/http/httptest"

	// "kitabisa-test/entity"
)

func TestSoccerGetAllTeamPositive(t *testing.T) {
	req, err := http.NewRequest("GET", "/soccer/team", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTeams)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSoccerGetTeamPositive(t *testing.T) {
	req, err := http.NewRequest("GET", "/soccer/team/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTeamDetail)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSoccerAddTeamPositive(t *testing.T) {
	var jsonStr = []byte(`{"name": "Barcelona FC"}`)

	req, err := http.NewRequest("POST", "/soccer/team", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddTeam)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSoccerRemoveTeamPositive(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/soccer/team/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RemoveTeam)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSoccerAddPlayerPositive(t *testing.T) {
	var jsonStr = []byte(`{"name": "Ronaldo"}`)

	req, err := http.NewRequest("POST", "/soccer/player", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddPlayer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSoccerRemovePlayerPositive(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/soccer/player/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RemovePlayer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}