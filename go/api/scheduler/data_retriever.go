package scheduler

import (
	"api/data"
	"api/db"
	"api/security"
	cu "lib/collectionutils"
	"time"
)

type SchedulerData struct {
	Users           data.Users      `json:"users"`
	Teams           []*TeamInfo     `json:"teams"`
	Buildings       []*BuildingInfo `json:"buildings"`
	Rooms           []*RoomInfo     `json:"rooms"`
	Resources       data.Resources  `json:"resources"`
	CurrentBookings *data.Bookings  `json:"current_bookings"`
	PastBookings    *data.Bookings  `json:"past_bookings"`
	StartDate       *time.Time      `json:"start_date"`
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

	// Get past weeks bookings
	weekAgoFrom := from.Add(-1 * time.Hour * 24 * 7) // subtract a week
	weekAgoTo := to.Add(-1 * time.Hour * 24 * 7)     // subtract a week
	pastBookingsInfo, err := GetBookings(weekAgoFrom, weekAgoTo)
	if err != nil {
		return nil, err
	}

	// schedulerData := SchedulerData{
	// 	Users:     users,
	// 	Teams:     teams,
	// 	Buildings: buildings,
	// 	Rooms:     rooms,
	// 	Resources: resources,
	// 	Bookings:  bookingsInfo,
	// }

	schedulerData := SchedulerData{
		Users:           users,
		Teams:           teams,
		Buildings:       buildings,
		Rooms:           rooms,
		Resources:       resources,
		CurrentBookings: &bookingsInfo.Bookings,
		PastBookings:    &pastBookingsInfo.Bookings,
		StartDate:       &from,
	}

	return &schedulerData, nil
}

// Group by building will partition the schedulerData into data that is only relevant to certain buildings
// this includes resources in a certain building, as well as users and bookings associated with a certain building
// Users in teams will also be seperated***
// Additionally, any data not associated with a specific building will be placed under the key ""
func GroupByBuilding(schedulerData *SchedulerData) map[string]*SchedulerData { // map[building_id]SchedulerData
	groupedData := make(map[string]*SchedulerData) // map[building_id]SchedulerData

	///////////////////
	// Create maps

	// Resources map
	resourcesMap := make(map[string]*data.Resource)
	for _, resource := range schedulerData.Resources {
		resourcesMap[*resource.Id] = resource
	}

	// Users map
	usersMap := make(map[string]*data.User)
	for _, user := range schedulerData.Users {
		usersMap[*user.Id] = user
	}

	// Room map
	roomsMap := make(map[string]*RoomInfo)
	for _, room := range schedulerData.Rooms {
		roomsMap[*room.Id] = room
	}

	////////////////////////
	// Group base objects

	// Group buildings by their id
	_, groupedBuildings := cu.GroupBy(schedulerData.Buildings, func(building *BuildingInfo) string {
		return *building.Id
	})

	groupUser := func(user *data.User) string {
		if user.BuildingID == nil {
			return ""
		}
		return *user.BuildingID
	}

	// Group users
	_, groupedUsers := cu.GroupBy(schedulerData.Users, groupUser)

	groupRoom := func(room *RoomInfo) string {
		if room.BuildingId == nil {
			return ""
		}
		return *room.BuildingId
	}
	// Group rooms
	_, groupedRooms := cu.GroupBy(schedulerData.Rooms, groupRoom)

	groupResource := func(resource *data.Resource) string {
		if roomsMap[*resource.RoomId].BuildingId == nil {
			return ""
		}
		return *roomsMap[*resource.RoomId].BuildingId
	}

	// Group resources
	_, groupedResources := cu.GroupBy(schedulerData.Resources, groupResource)

	// If bookings already had a resource assigned, group them by that resource,
	// otherwise, group them by the assigned users group
	groupBooking := func(booking *data.Booking) string {
		if booking.ResourceId != nil {
			return groupResource(resourcesMap[*booking.ResourceId])
		}
		return groupUser(usersMap[*booking.UserId])
	}

	// Group current bookings
	_, groupedCurrentBookings := cu.GroupBy(*schedulerData.CurrentBookings, groupBooking)
	// Group past bookings
	_, groupedPastBookings := cu.GroupBy(*schedulerData.PastBookings, groupBooking)

	////////////////////////////////////////
	// Group sub arrays inside infos objects
	groupedTeamUserIds := make(map[string](map[string][]string))
	for _, team := range schedulerData.Teams {
		_, groupedTeamUserIds[*team.Id] = cu.GroupBy(team.UserIds, func(userId string) string {
			return groupUser(usersMap[userId])
		})
	}

	groupedBuildingRoomIds := make(map[string](map[string][]string))
	for _, building := range schedulerData.Buildings {
		_, groupedBuildingRoomIds[*building.Id] = cu.GroupBy(building.RoomIds, func(roomId string) string {
			return groupRoom(roomsMap[roomId])
		})
	}

	groupedRoomResourceIds := make(map[string](map[string][]string))
	for _, room := range schedulerData.Rooms {
		_, groupedRoomResourceIds[*room.Id] = cu.GroupBy(room.ResourceIds, func(resourceId string) string {
			return groupResource(resourcesMap[resourceId])
		})
	}

	/////////////////////////////////////////////
	// Create new grouped scheduler data objects

	// Assign buildings
	for buildingId, buildings := range groupedBuildings {
		groupedData[buildingId] = &SchedulerData{
			Buildings:       buildings,
			Users:           data.Users{},
			Teams:           []*TeamInfo{},
			Rooms:           []*RoomInfo{},
			Resources:       data.Resources{},
			CurrentBookings: &data.Bookings{},
			PastBookings:    &data.Bookings{},
			StartDate:       schedulerData.StartDate,
		}

		// newBuildings := []*BuildingInfo{}
		// for index, buildingInfo := range schedulerData.Buildings {
		// 	newBuildings = append(newBuildings, &BuildingInfo{
		// 		buildingInfo.Building,
		// 		groupedBuildingRoomIds[buildingId][buildingId],
		// 	})
		// 	if groupedBuildingRoomIds[buildingId][buildingId] == nil {
		// 		newBuildings[index].RoomIds = []string{}
		// 	}
		// }

		// Add the teams
		newTeams := []*TeamInfo{}
		for index, teamInfo := range schedulerData.Teams {
			newTeams = append(newTeams, &TeamInfo{
				teamInfo.Team,
				groupedTeamUserIds[*teamInfo.Id][buildingId],
			})
			if groupedTeamUserIds[*teamInfo.Id][buildingId] == nil {
				newTeams[index].UserIds = []string{}
			}
		}
		groupedData[buildingId].Teams = newTeams

		groupedData[buildingId].Users = groupedUsers[buildingId]

		newRooms := []*RoomInfo{}
		for index, roomInfo := range groupedRooms[buildingId] {
			newRooms = append(newRooms, &RoomInfo{
				roomInfo.Room,
				groupedRoomResourceIds[*roomInfo.Id][buildingId],
			})
			if groupedRoomResourceIds[*roomInfo.Id][buildingId] == nil {
				newRooms[index].ResourceIds = []string{}
			}
		}
		groupedData[buildingId].Rooms = groupedRooms[buildingId]
		// Correct

		currBookings := groupedCurrentBookings[buildingId]
		groupedData[buildingId].CurrentBookings = (*data.Bookings)(&currBookings)

		pastBookings := groupedPastBookings[buildingId]
		groupedData[buildingId].PastBookings = (*data.Bookings)(&pastBookings)

		resources := groupedResources[buildingId]
		groupedData[buildingId].Resources = resources
	}

	return groupedData
}
