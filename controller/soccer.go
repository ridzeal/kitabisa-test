package controller

import (
	"fmt"
	"gorm.io/gorm"
	"kitabisa-test/entity"
	"kitabisa-test/models"
)

// GetTeams fetching all teams in databases
func GetTeams(db *gorm.DB) (teams []entity.Team, err error) {
	var teamResult []models.Team
	db.Find(&teamResult)
	fmt.Printf("Result %v", teamResult)

	for _, val := range(teamResult) {
		teams = append(teams, entity.Team{TeamID: val.ID, TeamName: val.Name})
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