package handler

import (
	"net/http"
	"fmt"

	"kitabisa-test/entity"
)

// GetTeams returns list of soccer team
func GetTeams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app := ctx.Value(entity.AppCtx)
	fmt.Printf("%v", app)

	response := &entity.GetTeamResponse{
		Teams: []string{"test","test2"},
	}
	renderResponse(w, r, http.StatusOK, response)
}

// GetTeamDetail returns list of players in a team
func GetTeamDetail(w http.ResponseWriter, r *http.Request) {
	response := &entity.GetTeamDetailResponse{
		Players: []string{"test","test2"},
	}
	renderResponse(w, r, http.StatusOK, response)
}

// AddTeam to database
func AddTeam(w http.ResponseWriter, r *http.Request) {
	response := &entity.AddTeamResponse{
		TeamID: "autogen",
		Status: "200",
		Error: "",
	}
	renderResponse(w, r, http.StatusOK, response)
}

// RemoveTeam from database
func RemoveTeam(w http.ResponseWriter, r *http.Request) {
	response := &entity.RemoveTeamResponse{
		Status: "200",
		Error: "",
	}
	renderResponse(w, r, http.StatusOK, response)
}

// AddPlayer to a team
func AddPlayer(w http.ResponseWriter, r *http.Request) {
	response := &entity.AddPlayerResponse{
		PlayerID: "autogen",
		Status: "200",
		Error: "",
	}
	renderResponse(w, r, http.StatusOK, response)
}

// RemovePlayer from a team
func RemovePlayer(w http.ResponseWriter, r *http.Request) {
	response := &entity.RemovePlayerResponse{
		Status: "200",
		Error: "",
	}
	renderResponse(w, r, http.StatusOK, response)
}