package controller

import (
	"gorm.io/gorm"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSoccerGetTeamPositive(t *testing.T) {
	db := gorm.DB{}
	teams, err := GetTeams(&db)
	assert.NotEqual(t, 0, len(teams), "Team list shouldn't empty")
	assert.Equal(t, nil, err, "There should be no error")
}

func TestSoccerGetTeamDetailPositive(t *testing.T) {
	teamID := 1
	team, err := GetTeamDetail(teamID)
	assert.NotEqual(t, 0, team.TeamID, "On success, team ID should be returned and not 0")
	assert.Equal(t, nil, err, "There should be no error")
}

func TestSoccerAddTeamPositive(t *testing.T) {
	teamName := "Real Madrid"
	teamID, err := AddTeam(teamName)
	assert.NotEqual(t, 0, teamID, "On success, team ID should be returned and not 0")
	assert.Equal(t, nil, err, "There should be no error")
}

func TestSoccerRemoveTeamPositive(t *testing.T) {
	teamID := 1
	err := RemoveTeam(teamID)
	assert.Equal(t, nil, err, "There should be no error")
}