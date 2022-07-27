package scheduler

import (
	"api/data"
	"api/db"
	"api/security"
	"time"
)

type SchedulerData struct {
	Users     data.Users      `json:"users"`
	Teams     []*TeamInfo     `json:"teams"`
	Buildings []*BuildingInfo `json:"buildings"`
	Rooms     []*RoomInfo     `json:"rooms"`
	Resources data.Resources  `json:"resources"`
	Bookings  *BookingInfo    `json:"bookings"`
}

type BookingInfo struct {
	From     time.Time     `json:"from"`
	To       time.Time     `json:"to"`
	Bookings data.Bookings `json:"bookings"`
}

type TeamInfo struct {
	*data.Team
	UserIds []string `json:"user_ids"`
}

type BuildingInfo struct {
	*data.Building
	RoomIds []string `json:"room_ids"`
}

type RoomInfo struct {
	*data.Room
	ResourceIds []string `json:"resource_ids"`
}

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
func GetBookings(from time.Time, to time.Time) (*BookingInfo, error) {
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")}
	// Connect to db
	access, err := db.Open()
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
	bookingInfo := BookingInfo{
		From:     from,
		To:       to,
		Bookings: bookings,
	}
	return &bookingInfo, nil
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
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "TEAM", "USER"),
		data.CreateGenericPermission("VIEW", "USER", "TEAM")}
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
func GetCompiledTeamInfo() ([]*TeamInfo, error) {
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
		teamInfosMap[*team.Id] = &TeamInfo{team, []string{}}
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

// GetBuildings retrieves all the buildings from the database
func GetBuildings() (data.Buildings, error) {
	// Create a database connection
	access, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer access.Close()

	da := data.NewResourceDA(access)
	building := data.Building{}
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "BUILDING")}
	buildings, err := da.FindBuildingResource(&building, permissions)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		return nil, err
	}
	return buildings, nil
}

// IMPORTANT: It is assumed at this point that rooms are flat

// TODO: @JonathanEnslin Get funcs for resources and rooms, similar to teams/users
// GetRooms retrieves all the rooms from the database
func GetRooms() (data.Rooms, error) {
	// Create a database connection
	access, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer access.Close()

	da := data.NewResourceDA(access)
	room := data.Room{}
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "ROOM")}
	rooms, err := da.FindRoomResource(&room, permissions)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

// GetResourceIdentifiers retrieves all the Resources from the database
func GetResourceIdentifiers() (data.Resources, error) {
	// Create db connection
	access, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer access.Close()

	da := data.NewResourceDA(access)
	identifier := data.Resource{}
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "IDENTIFIER")}
	identifiers, err := da.FindIdentifier(&identifier, permissions)
	if err != nil {
		return nil, err
	}

	err = access.Commit()
	if err != nil {
		return nil, err
	}
	return identifiers, nil
}

// GetCompiledResourceInfo maps room ids to buildings, and resource ids to rooms
func GetCompileResourceInfo() ([]*BuildingInfo, []*RoomInfo, data.Resources, error) {
	buildings, err := GetBuildings()
	if err != nil {
		return nil, nil, nil, err
	}
	rooms, err := GetRooms()
	if err != nil {
		return nil, nil, nil, err
	}
	identifiers, err := GetResourceIdentifiers()
	if err != nil {
		return nil, nil, nil, err
	}

	// Map rooms to buildings
	buildingsMap := make(map[string]*BuildingInfo)
	for _, building := range buildings {
		buildingsMap[*building.Id] = &BuildingInfo{building, []string{}}
	}
	for _, room := range rooms {
		buildingsMap[*room.BuildingId].RoomIds = append(buildingsMap[*room.BuildingId].RoomIds, *room.Id)
	}

	buildingInfos := []*BuildingInfo{}
	for _, buildingsInfo := range buildingsMap {
		buildingInfos = append(buildingInfos, buildingsInfo)
	}

	// Map resources to rooms
	roomsMap := make(map[string]*RoomInfo)
	for _, room := range rooms {
		roomsMap[*room.Id] = &RoomInfo{room, []string{}}
	}
	for _, identifier := range identifiers {
		roomsMap[*identifier.RoomId].ResourceIds = append(roomsMap[*identifier.RoomId].ResourceIds, *identifier.Id)
	}

	roomInfos := []*RoomInfo{}
	for _, roomInfo := range roomsMap {
		roomInfos = append(roomInfos, roomInfo)
	}

	return buildingInfos, roomInfos, identifiers, nil
}

// GetSchedulerData retrieves all the data needed by the scheduler
func GetSchedulerData(from time.Time, to time.Time) (*SchedulerData, error) {
	// Get the users
	users, err := GetUsers()
	if err != nil {
		return nil, err
	}

	// Get the teams
	teams, err := GetCompiledTeamInfo()
	if err != nil {
		return nil, err
	}

	// Get the buildings, rooms, resources
	buildings, rooms, resources, err := GetCompileResourceInfo()
	if err != nil {
		return nil, err
	}

	// Get the existing bookings
	bookingsInfo, err := GetBookings(from, to)
	if err != nil {
		return nil, err
	}

	schedulerData := SchedulerData{
		Users:     users,
		Teams:     teams,
		Buildings: buildings,
		Rooms:     rooms,
		Resources: resources,
		Bookings:  bookingsInfo,
	}

	return &schedulerData, nil
}
