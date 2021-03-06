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
	ctx := r.Context()
	app, ok := ctx.Value(entity.AppCtx).(*entity.Application)
	if !ok {
		status := http.StatusUnprocessableEntity
    http.Error(w, http.StatusText(status), status)
    return
	}

	var addTeamRequest *entity.AddTeamRequest

	if err := parseBody(r.Body, r.Header, &addTeamRequest); err != nil {
		respBody := &entity.AddTeamResponse{
			TeamID: 0,
			Status: http.StatusBadRequest,
			Error: "Missing parameter",
		}

		renderResponse(w, r, http.StatusBadRequest, respBody)
		return
	}

	teamID, err := controller.AddTeam(app.DB, addTeamRequest.TeamName)
	if err != nil {
		respBody := &entity.AddTeamResponse{
			TeamID: 0,
			Status: http.StatusInternalServerError,
			Error: err.Error(),
		}

		renderResponse(w, r, http.StatusInternalServerError, respBody)
		return
	}

	response := &entity.AddTeamResponse{
		TeamID: teamID,
		Status: http.StatusOK,
		Error: "",
	}
	renderResponse(w, r, http.StatusOK, response)
}

// RemoveTeam from database
func RemoveTeam(w http.ResponseWriter, r *http.Request) {
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

	controller.RemoveTeam(app.DB, idInt)
	
	response := &entity.RemoveTeamResponse{
		Status: http.StatusOK,
		Error: "",
	}
	renderResponse(w, r, http.StatusOK, response)
}

// AddPlayer to a team
func AddPlayer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app, ok := ctx.Value(entity.AppCtx).(*entity.Application)
	if !ok {
		status := http.StatusUnprocessableEntity
    http.Error(w, http.StatusText(status), status)
    return
	}

	var addPlayerRequest *entity.AddPlayerRequest

	if err := parseBody(r.Body, r.Header, &addPlayerRequest); err != nil {
		respBody := &entity.AddPlayerResponse{
			PlayerID: 0,
			Status: http.StatusBadRequest,
			Error: "Missing parameter",
		}

		renderResponse(w, r, http.StatusBadRequest, respBody)
		return
	}

	// Check if Team is exist
	_, err := controller.GetTeamDetail(app.DB, addPlayerRequest.TeamID)
	if err != nil {
		respBody := &entity.AddPlayerResponse{
			PlayerID: 0,
			Status: http.StatusNotFound,
			Error: err.Error(),
		}

		renderResponse(w, r, http.StatusNotFound, respBody)
		return
	}

	playerID, err := controller.AddPlayer(
		app.DB,
		addPlayerRequest.PlayerName,
		addPlayerRequest.TeamID,
	)
	if err != nil {
		respBody := &entity.AddPlayerResponse{
			PlayerID: 0,
			Status: http.StatusInternalServerError,
			Error: err.Error(),
		}

		renderResponse(w, r, http.StatusInternalServerError, respBody)
		return
	}

	response := &entity.AddPlayerResponse{
		PlayerID: playerID,
		Status: 200,
		Error: "",
	}
	renderResponse(w, r, http.StatusOK, response)
}

// RemovePlayer from a team
func RemovePlayer(w http.ResponseWriter, r *http.Request) {
	playerID := chi.URLParam(r, "playerID")
	
	ctx := r.Context()
	app, ok := ctx.Value(entity.AppCtx).(*entity.Application)
	if !ok {
		status := http.StatusUnprocessableEntity
    http.Error(w, http.StatusText(status), status)
    return
	}

	idInt, err := strconv.Atoi(playerID)
	if err != nil {
		status := http.StatusBadRequest
    http.Error(w, http.StatusText(status), status)
    return
	}

	controller.RemovePlayer(app.DB, idInt)
	
	response := &entity.RemovePlayerResponse{
		Status: http.StatusOK,
		Error: "",
	}
	renderResponse(w, r, http.StatusOK, response)
}