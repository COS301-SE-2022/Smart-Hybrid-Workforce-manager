package data

import (
	"encoding/json"
	"lib/collectionutils"
	"lib/utils"
	"time"
)

type SchedulerConfig struct {
	DailyConfig  *Config `json:"daily_config"`
	WeeklyConfig *Config `json:"weekly_config"`
}

type Config struct {
	Seed             int     `json:"seed"`
	PopulationSize   int     `json:"populationSize"`
	Generations      int     `json:"generations"`
	MutationRate     float64 `json:"mutationRate"`
	CrossOverRate    float64 `json:"crossOverRate"`
	TournamentSize   int     `json:"tournamentSize"`
	TimeLimitSeconds int     `json:"time_limit_seconds"`
}

type SchedulerData struct {
	Users               Users               `json:"users"`
	Teams               []*TeamInfo         `json:"teams"`
	Roles               []*RoleInfo         `json:"roles"`
	Buildings           []*BuildingInfo     `json:"buildings"`
	Rooms               []*RoomInfo         `json:"rooms"`
	Resources           Resources           `json:"resources"`
	CurrentBookings     *Bookings           `json:"current_bookings"`
	PastBookings        *Bookings           `json:"past_bookings"`
	StartDate           *time.Time          `json:"start_date"`
	MeetingRoomBookings MeetingRoomBookings `json:"meeting_room_bookings"`
	// map[teamId]*TeamInfo
	TeamsMap map[string]*TeamInfo `json:"-"`

	// map[roleId]*RoleInfo
	RolesMap map[string]*RoleInfo `json:"-"`

	// map[roomId]*RoomInfo
	RoomsMap map[string]*RoomInfo `json:"-"`

	// map[resourceId]*Resource
	ResourcesMap map[string]*Resource `json:"-"`

	// map[UserId]*User
	UserMap map[string]*User `json:"-"`
}

type RoleInfo struct {
	*Role
	UserIds []string `json:"user_ids"`
}

type BookingInfo struct {
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
	Bookings Bookings  `json:"bookings"`
}

type TeamInfo struct {
	*Team
	UserIds []string `json:"user_ids"`
}

type BuildingInfo struct {
	*Building
	RoomIds []string `json:"room_ids"`
}

type RoomInfo struct {
	*Room
	ResourceIds []string `json:"resource_ids"`
}

// User identifies a user via common attributes
type User struct {
	Id                 *string    `json:"id,omitempty"`
	Identifier         *string    `json:"identifier,omitempty"`
	FirstName          *string    `json:"first_name,omitempty"`
	LastName           *string    `json:"last_name,omitempty"`
	Email              *string    `json:"email,omitempty"`
	Picture            *string    `json:"picture,omitempty"`
	DateCreated        time.Time  `json:"date_created,omitempty"`
	WorkFromHome       *bool      `json:"work_from_home,omitempty"`
	Parking            *string    `json:"parking,omitempty"`
	OfficeDays         *int       `json:"office_days,omitempty"`
	PreferredStartTime *time.Time `json:"preferred_start_time,omitempty"`
	PreferredEndTime   *time.Time `json:"preferred_end_time,omitempty"`
	PreferredDesk      *string    `json:"preferred_desk,omitempty"`
}

// Users represent a splice of User
type Users []*User

// Resource identifies a Resource via common attributes
type Resource struct {
	Id                      *string  `json:"id,omitempty"`
	RoomId                  *string  `json:"room_id,omitempty"`
	Name                    *string  `json:"name,omitempty"`
	XCoord                  *float64 `json:"xcoord,omitempty"`
	YCoord                  *float64 `json:"ycoord,omitempty"`
	Width                   *float64 `json:"width,omitempty"`
	Height                  *float64 `json:"height,omitempty"`
	Rotation                *float64 `json:"rotation,omitempty"`
	RoleId                  *string  `json:"role_id,omitempty"`
	ResourceType            *string  `json:"resource_type,omitempty"`
	Decorations             *string  `json:"decorations,omitempty"`
	DateCreated             *string  `json:"date_created,omitempty"`
	cachedParsedDecorations *map[string]any
	atCachedDecorationsPtr  *string
}

// Resources represent a splice of Resource
type Resources []*Resource

// Booking identifies a booking via common attributes
type Booking struct {
	Id                   *string    `json:"id,omitempty"`
	UserId               *string    `json:"user_id,omitempty"`
	ResourceType         *string    `json:"resource_type,omitempty"`
	ResourcePreferenceId *string    `json:"resource_preference_id,omitempty"`
	ResourceId           *string    `json:"resource_id,omitempty"`
	Start                *time.Time `json:"start,omitempty"`
	End                  *time.Time `json:"end,omitempty"`
	Booked               *bool      `json:"booked,omitempty"`
	Automated            *bool      `json:"automated,omitempty"`
	Dependent            *string    `json:"dependent,omitempty"`
	DateCreated          *time.Time `json:"date_created,omitempty"`
}

// MeetingRoomBooking identifies a meeting room booking
type MeetingRoomBooking struct {
	Booking                  *Booking `json:"booking,omitempty"` // Consider removing, two api calls for meeting bookings
	Id                       *string  `json:"id,omitempty"`
	BookingId                *string  `json:"booking_id,omitempty"`
	TeamId                   *string  `json:"team_id,omitempty"`
	RoleId                   *string  `json:"role_id,omitempty"`
	AdditionalAttendees      *int     `json:"additional_attendees,omitempty"`
	DesksAttendees           *bool    `json:"desks_attendees,omitempty"`
	DesksAdditionalAttendees *bool    `json:"desks_additional_attendees,omitempty"`
}

type MeetingRoomBookings []*MeetingRoomBooking

// Bookings represent a splice of Booking
type Bookings []*Booking

// Identifier identifies a user via common attributes
type Team struct {
	Id          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Capacity    *int       `json:"capacity,omitempty"`
	Picture     *string    `json:"picture,omitempty"`
	Priority    *int       `json:"priority,omitempty"`
	TeamLeadId  *string    `json:"team_lead_id,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
}

type Role struct {
	Id        *string    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Color     *string    `json:"color,omitempty"`
	LeadId    *string    `json:"lead_id,omitempty"`
	DateAdded *time.Time `json:"date_added,omitempty"`
}

// Building identifies a Building Resource via common attributes
type Building struct {
	Id        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	Location  *string `json:"location,omitempty"`
	Dimension *string `json:"dimension,omitempty"`
}

// Room identifies a Room Resource via common attributes
type Room struct {
	Id         *string  `json:"id,omitempty"`
	BuildingId *string  `json:"building_id,omitempty"`
	Name       *string  `json:"name,omitempty"`
	XCoord     *float64 `json:"xcoord,omitempty"`
	YCoord     *float64 `json:"ycoord,omitempty"`
	ZCoord     *float64 `json:"zcoord,omitempty"`
	Dimension  *string  `json:"dimension,omitempty"`
}

// Parses resource decorations
func (r *Resource) GetDecorations() map[string]any {
	if r.cachedParsedDecorations != nil && r.atCachedDecorationsPtr == r.Decorations {
		return *r.cachedParsedDecorations
	}
	r.atCachedDecorationsPtr = r.Decorations
	parsed := map[string]any{}
	if r.Decorations == nil {
		return parsed
	}
	err := json.Unmarshal([]byte(*r.Decorations), &parsed)
	if err != nil {
		return map[string]any{}
	}
	r.cachedParsedDecorations = &parsed
	return parsed
}

// Gets the capacity (most likely of meeting rooms)
func (r *Resource) GetCapacity() int {
	if collectionutils.MapHasKey(r.GetDecorations(), "capacity") {
		if num, ok := r.GetDecorations()["capacity"].(float64); ok {
			return int(num)
		}
	}
	return -1
}

// =======================
// ====    Methods    ====
// =======================
func (b *Booking) GetWeekday() time.Weekday {
	return (b.Start.Weekday() + 6) % 7
}

// ExtractUserIdsDuplicates creates an array of user ids which are duplicated
func ExtractUserIdsDuplicates(schedulerData *SchedulerData) []string {
	// Keep track of how many days users are already coming into the office
	timesAlreadyComingIn := make(map[string]int, 0) // (map[user id]times coming in already)
	for _, booking := range *schedulerData.CurrentBookings {
		timesAlreadyComingIn[*booking.UserId]++ // Add one to indicate they are coming in
	}

	// Add users as many times as they need to come into office
	usersToAdd := []string{}
	for _, user := range schedulerData.Users {
		// Add them times they have to come in - days they already come in
		for i := 0; i < *user.OfficeDays-timesAlreadyComingIn[*user.Id]; i++ {
			usersToAdd = append(usersToAdd, *user.Id)
		}
	}
	return usersToAdd
}

// ExtractResourceIds extracts all resource ids from schedulerdata into a string array
func ExtractResourceIds(schedulerData *SchedulerData) []string {
	// Add resources to string array
	resources := []string{}
	for _, resource := range schedulerData.Resources {
		resources = append(resources, *resource.Id)
	}
	return resources
}

// ExtractAvailable deskids extracts all available desk ids from schedulerdata into a string array
func ExtractAvailableDeskIds(schedulerData *SchedulerData) []string {
	// Build map of all resources that have arleady been assigned to a booking
	bookedResources := make(map[string]bool)
	for _, booking := range *schedulerData.CurrentBookings {
		if booking.ResourceId != nil && *booking.ResourceId != "" {
			bookedResources[*booking.ResourceId] = true
		}
	}

	// Add resources to string array
	resources := []string{}
	for _, resource := range schedulerData.Resources {
		if resource.ResourceType != nil &&
			*resource.ResourceType == "DESK" &&
			!collectionutils.MapHasKey(bookedResources, *resource.Id) { // Check if resource is a desk
			resources = append(resources, *resource.Id) // Append to available resources
		}
	}
	return resources
}

// ExtractUserIdMap extracts all user ids from the schedulerdata into an indexed map
func ExtractUserIdMap(schedulerData *SchedulerData) map[int](string) {
	var result map[int](string)
	result = make(map[int]string)
	for i, booking := range *schedulerData.CurrentBookings {
		result[i] = *booking.UserId
	}
	return result
}

// ExtractUserIdMap extracts all user ids from the schedulerdata into an indexed map
// capped by the amount of resources that are available, this is incase resources are deleted
// between weekly scheduler running and daily scheduler called
func ExtractUserIdMapAsMuchAsAvailable(schedulerData *SchedulerData) map[int](string) {
	result := make(map[int]string) // map[index]userId
	bookingIndices := collectionutils.SequentialSequence(0, len(*schedulerData.CurrentBookings))
	// shuffle the indices
	for i := range bookingIndices {
		j := utils.RandInt(0, i+1)
		bookingIndices[i], bookingIndices[j] = bookingIndices[j], bookingIndices[i]
	}
	// For the number of available resources, extract user ids from the bookings according to the random indices
	for i := range schedulerData.Resources {
		if i >= len((*schedulerData.CurrentBookings)) {
			break
		}
		result[i] = *(*schedulerData.CurrentBookings)[bookingIndices[i]].UserId
	}
	return result
}

// ExtractMecessaryUserMap extracts all user ids from the schedulerdata into an indexed map
// capped by the amount of resources that are available, this is incase resources are deleted
// between weekly scheduler running and daily scheduler called, and given that their booking
// does not already have an assigned resource
func ExtractNecessaryUserMap(schedulerData *SchedulerData) map[int](string) {
	result := make(map[int]string) // map[index]userId
	availableDesks := ExtractAvailableDeskIds(schedulerData)
	// shuffle the current bookings
	for i := range *schedulerData.CurrentBookings {
		j := utils.RandInt(0, i+1)
		(*schedulerData.CurrentBookings)[i], (*schedulerData.CurrentBookings)[j] = (*schedulerData.CurrentBookings)[j], (*schedulerData.CurrentBookings)[i]
	}
	// For the number of available resources, extract user ids from the bookings according to the random indices
	assignIndex := 0
	for i := 0; i < len(*schedulerData.CurrentBookings); i++ {
		if assignIndex >= len(availableDesks) {
			break
		}
		if (*schedulerData.CurrentBookings)[i].ResourceId == nil {
			result[assignIndex] = *(*schedulerData.CurrentBookings)[i].UserId
			assignIndex++
		}
	}
	return result
}

func ExtractInverseUserIdMap(userIdMap map[int]string) map[string][]int {
	userIdToIndicesMap := make(map[string][]int)
	for index, userId := range userIdMap {
		if !collectionutils.MapHasKey(userIdToIndicesMap, userId) {
			userIdToIndicesMap[userId] = make([]int, 0)
		}
		userIdToIndicesMap[userId] = append(userIdToIndicesMap[userId], index)
	}
	return userIdToIndicesMap
}

func (data *SchedulerData) MapRooms() {
	data.RoomsMap = map[string]*RoomInfo{}
	for _, roomInfo := range data.Rooms {
		data.RoomsMap[*roomInfo.Id] = roomInfo
	}
}

func (data *SchedulerData) MapResources() {
	data.ResourcesMap = map[string]*Resource{}
	for _, resource := range data.Resources {
		data.ResourcesMap[*resource.Id] = resource
	}
}

func (data *SchedulerData) MapUsers() {
	data.UserMap = map[string]*User{}
	for _, user := range data.Users {
		data.UserMap[*user.Id] = user
	}
}

func (data *SchedulerData) MapTeams() {
	data.TeamsMap = map[string]*TeamInfo{}
	for _, teamInfo := range data.Teams {
		data.TeamsMap[*teamInfo.Id] = teamInfo
	}
}

func (data *SchedulerData) MapRoles() {
	data.RolesMap = map[string]*RoleInfo{}
	if data.Roles != nil && len(data.Roles) == 0 {
		for _, roleInfo := range data.Roles {
			data.RolesMap[*roleInfo.Id] = roleInfo
		}
	}
}

func (data *SchedulerData) ApplyMapping() {
	data.MapRooms()
	data.MapResources()
	data.MapTeams()
	data.MapUsers()
	data.MapRoles()
}
