package data

import (
	"lib/collectionutils"
	"time"
)

type Config struct {
	Seed           int     `json:"seed"`
	PopulationSize int     `json:"populationSize"`
	Generations    int     `json:"generations"`
	MutationRate   float64 `json:"mutationRate"`
	CrossOverRate  float64 `json:"crossOverRate"`
	TournamentSize int     `json:"tournamentSize"`
}

type SchedulerData struct {
	Users           Users           `json:"users"`
	Teams           []*TeamInfo     `json:"teams"`
	Buildings       []*BuildingInfo `json:"buildings"`
	Rooms           []*RoomInfo     `json:"rooms"`
	Resources       Resources       `json:"resources"`
	CurrentBookings *Bookings       `json:"current_bookings"`
	PastBookings    *Bookings       `json:"past_bookings"`
	StartDate       *time.Time      `json:"start_date"`
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
	Id           *string  `json:"id,omitempty"`
	RoomId       *string  `json:"room_id,omitempty"`
	Name         *string  `json:"name,omitempty"`
	XCoord       *float64 `json:"xcoord,omitempty"`
	YCoord       *float64 `json:"ycoord,omitempty"`
	Width        *float64 `json:"width,omitempty"`
	Height       *float64 `json:"height,omitempty"`
	Rotation     *float64 `json:"rotation,omitempty"`
	RoleId       *string  `json:"role_id,omitempty"`
	ResourceType *string  `json:"resource_type,omitempty"`
	Decorations  *string  `json:"decorations,omitempty"`
	DateCreated  *string  `json:"date_created,omitempty"`
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
		for i := 0; i < *user.OfficeDays-timesAlreadyComingIn[*user.Id]; i++ { // TODO: @JonathanEnslin find out what do if no office days?
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

// ExtractResourceIds extracts all available desk ids from schedulerdata into a string array
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
