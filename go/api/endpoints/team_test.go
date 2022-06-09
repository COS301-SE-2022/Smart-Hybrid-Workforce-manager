package endpoints

import (
	"api/data"
	"testing"
)

func createTeam(id string, name string, description string, capacity int, picture string) data.Team {
	var team data.Team
	team.Id = &id
	team.Name = &name
	team.Description = &description
	team.Capacity = &capacity
	team.Picture = &picture
	return team
}

func TestCreateTeamHandler(t *testing.T) {

}
