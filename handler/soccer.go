package handler

import (
	"net/http"
	"strconv"
	"github.com/go-chi/chi"
	
	"kitabisa-test/entity"
	"kitabisa-test/controller"
)

// GetTeams returns list of soccer team
func GetTeams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	app, ok := ctx.Value(entity.AppCtx).(*entity.Application)
	if !ok {
		status := http.StatusUnprocessableEntity
    http.Error(w, http.StatusText(status), status)
    return
  }
	
	teams, err := controller.GetTeams(app.DB)
	if err != nil {
		status := http.StatusInternalServerError
		http.Error(w, http.StatusText(status), status)
	}

	response := &entity.GetTeamResponse{
		Teams: teams,
	}
	renderResponse(w, r, http.StatusOK, response)
}

// GetTeamDetail returns list of players in a team
func GetTeamDetail(w http.ResponseWriter, r *http.Request) {
	teamID := chi.URLParam(r, "teamID")
	
	ctx := r.Context()
	app, ok := ctx.Value(entity.AppCtx).(*entity.Application)
	if !ok {
		status := http.StatusUnprocessableEntity
    http.Error(w, http.StatusText(status), status)
    return
	}

	idInt, err := strconv.Atoi(teamID)
	if err != nil {
		status := http.StatusBadRequest
    http.Error(w, http.StatusText(status), status)
    return
	}

	team, err := controller.GetTeamDetail(app.DB, idInt)
	
	response := &entity.GetTeamDetailResponse{
		Team: team,
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