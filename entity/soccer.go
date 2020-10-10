package entity

// Team is an object that represents a team
type Team struct {
	TeamID int `json:"team_id"`
	TeamName string `json:"team_name"`
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
	Players []string `json:"players"`
}

// AddTeamResponse is response object after adding a team
type AddTeamResponse struct {
	TeamID string `json:"team_id"`
	Status string `json:"status"`
	Error string `json:"error"`
}

// RemoveTeamResponse is response object after removing a team
type RemoveTeamResponse struct {
	Status string `json:"status"`
	Error string `json:"error"`
}

// AddPlayerResponse is response object after adding a player into the team
type AddPlayerResponse struct {
	PlayerID string `json:"player_id"`
	Status string `json:"status"`
	Error string `json:"error"`
}

// RemovePlayerResponse is response object after removing a player from the team
type RemovePlayerResponse struct {
	Status string `json:"status"`
	Error string `json:"error"`
}