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
func GetTeamDetail(db *gorm.DB, teamID int) (team entity.TeamDetail, err error) {
	var teamResult models.Team
	var playerList []models.Player

	result := db.First(&teamResult, teamID)
	err = result.Error
	db.Where("team = ?", teamID).Find(&playerList)

	team = entity.TeamDetail{
		TeamID: teamResult.ID,
		TeamName: teamResult.Name,
		Players: []entity.Player{},
	}

	for _, val := range(playerList) {
		team.Players = append(team.Players, entity.Player{
			PlayerID: val.ID,
			PlayerName: val.Name,
		})
	}

	return team, err
}

// AddTeam to database
func AddTeam(db *gorm.DB, teamName string) (teamID int, err error) {
	team := models.Team{Name: teamName}
	result := db.Create(&team)
	teamID = team.ID
	err = result.Error
	return teamID, err
}

// RemoveTeam from database
func RemoveTeam(db *gorm.DB, teamID int) (err error) {
	db.Delete(&models.Team{}, teamID)
	return err
}

// AddPlayer to database
func AddPlayer(db *gorm.DB, playerName string, teamID int) (playerID int, err error) {
	player := models.Player{Name: playerName, Team: teamID}
	result := db.Create(&player)
	playerID = player.ID
	err = result.Error
	return playerID, err
}

// RemovePlayer from database
func RemovePlayer(db *gorm.DB, playerID int) (err error) {
	db.Delete(&models.Player{}, playerID)
	return err
}