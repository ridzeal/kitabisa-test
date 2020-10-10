package entity

// Team is an object that represents a team
type Team struct {
	TeamID int `json:"team_id"`
	TeamName string `json:"team_name"`
}

// TeamDetail shows team info and players
type TeamDetail struct {
	TeamID int `json:"team_id"`
	TeamName string `json:"team_name"`
	Players []Player `json:"players"`
}

// Player is an object that represents a player
type Player struct {
	PlayerID int `json:"player_id"`
	PlayerName string `json:"player_name"`
}

// GetTeamResponse is response object when get team API called
type GetTeamResponse struct {
	Teams []Team `json:"teams"`
}

// GetTeamDetailResponse is response object when get player API called
type GetTeamDetailResponse struct {
	Team TeamDetail `json:"team"`
}

// AddTeamRequest format when call API add team
type AddTeamRequest struct {
	TeamName string `json:"team_name"`
}

// AddTeamResponse is response object after adding a team
type AddTeamResponse struct {
	TeamID int `json:"team_id"`
	Status int `json:"status"`
	Error string `json:"error"`
}

// RemoveTeamResponse is response object after removing a team
type RemoveTeamResponse struct {
	Status int `json:"status"`
	Error string `json:"error"`
}

// AddPlayerRequest format when call API add player
type AddPlayerRequest struct {
	PlayerName string `json:"player_name"`
	TeamID int `json:"team_id"`
}

// AddPlayerResponse is response object after adding a player into the team
type AddPlayerResponse struct {
	PlayerID int `json:"player_id"`
	Status int `json:"status"`
	Error string `json:"error"`
}

// RemovePlayerResponse is response object after removing a player from the team
type RemovePlayerResponse struct {
	Status string `json:"status"`
	Error string `json:"error"`
}