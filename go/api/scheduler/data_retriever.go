package scheduler

import (
	"api/data"
	"api/db"
	"api/security"
	"time"
)

type TeamInfo struct {
	Id      *string    `json:"id"`
	Team    *data.Team `json:"team"`
	UserIds []string   `json:"user_ids"`
}

type TeamInfos []*TeamInfo

// GetUsers Retrieves all the users from the database
func GetUsers() (data.Users, error) {
	access, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer access.Close()

	da := data.NewUserDA(access)

	user := data.User{}
	users, err := da.FindIdentifier(&user)
	if err != nil {
		return nil, err
	}

	err = access.Commit()
	if err != nil {
		return nil, err
	}
	// TODO: @JonathanEnslin remove unnecessary fields from users
	return users, nil
}

// GetBookings Retrieves all the bookings from the database between from from to to
func GetBookings(from time.Time, to time.Time) (data.Bookings, error) {
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")}
	// Connect to db
	access, err := db.Open() // TODO: @JonathanEnslin
	if err != nil {
		return nil, err
	}
	defer access.Close()

	bookingFilter := data.Booking{
		Start: &from,
		End:   &to,
	}

	da := data.NewBookingDA(access)
	bookings, err := da.FindIdentifier(&bookingFilter, security.RemoveRolePermissions(permissions))
	if err != nil {
		return nil, err
	}

	// commit transaction
	err = access.Commit()
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

// GetTeams retrieves all the teams from the database
// IMPORTANT: At this point teams are assumed to be flat
func GetTeams() (data.Teams, error) {
	// Create a database connection
	access, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer access.Close()

	da := data.NewTeamDA(access)
	team := data.Team{}
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "TEAM", "IDENTIFIER")}
	teams, err := da.FindIdentifier(&team, permissions)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		return nil, err
	}
	return teams, nil
}

// GetUserTeams will return all the userteam pairs
func GetUserTeams() (data.UserTeams, error) {
	// Create a database connection
	access, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer access.Close()
	da := data.NewTeamDA(access)

	userTeam := data.UserTeam{}
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "TEAM", "USER")}
	userTeams, err := da.FindUserTeam(&userTeam, permissions)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		return nil, err
	}
	return userTeams, nil
}

// GetCompiledTeamInfo will return an array, with team_id as keys
// this map will contain the info of the team, as well as all the users belonging to
// that team
func GetCompiledTeamInfo() (TeamInfos, error) {
	teams, err := GetTeams()
	if err != nil {
		return nil, err
	}
	userTeams, err := GetUserTeams()
	if err != nil {
		return nil, err
	}
	teamInfosMap := make(map[string]*TeamInfo)
	for _, team := range teams {
		teamInfosMap[*team.Id] = &TeamInfo{Id: team.Id, Team: team, UserIds: []string{}}
	}
	for _, userTeam := range userTeams {
		teamInfosMap[*userTeam.TeamId].UserIds = append(teamInfosMap[*userTeam.TeamId].UserIds, *userTeam.UserId)
	}
	teamInfos := make([]*TeamInfo, len(teamInfosMap))
	i := 0
	for _, v := range teamInfosMap {
		teamInfos[i] = v
		i++
	}

	// final result would be for example:
	// [{"team_id": "122-122", "team": {team_info...}, "user_ids": ["123-123", "321-123"]},...]
	return teamInfos, nil
}
