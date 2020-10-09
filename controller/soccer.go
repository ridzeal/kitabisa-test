package controller

import (
	"kitabisa-test/entity"
)

// GetTeams fetching all teams in databases
func GetTeams() (teams []entity.Team, err error) {
	teams = []entity.Team{
		{TeamID: 1, TeamName: "Test"},
	}

	return teams, err
}

// GetTeamDetail fetching a team detail from database
func GetTeamDetail(teamID int) (team entity.Team, err error) {
	team = entity.Team{
		TeamID: 1, TeamName: "Test2",
	}

	return team, err
}

// AddTeam to database
func AddTeam(teamName string) (teamID int, err error) {
	teamID = 1
	return teamID, err
}

// RemoveTeam from database
func RemoveTeam(teamID int) (err error) {
	return err
}