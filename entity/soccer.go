package entity

// GetTeamResponse is response object when get team API called
type GetTeamResponse struct {
	Teams []string `json:"teams"`
}

// GetTeamDetailResponse is response object when get player API called
type GetTeamDetailResponse struct {
	Players []string `json:"players"`
}

// AddTeamResponse is response object after adding a team
type AddTeamResponse struct {
	TeamID string `json:"teamId"`
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
	PlayerID string `json:"playerId"`
	Status string `json:"status"`
	Error string `json:"error"`
}

// RemovePlayerResponse is response object after removing a player from the team
type RemovePlayerResponse struct {
	Status string `json:"status"`
	Error string `json:"error"`
}