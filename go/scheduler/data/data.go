package data

import "time"

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
	Id         *string `json:"id,omitempty"`
	BuildingId *string `json:"building_id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Location   *string `json:"location,omitempty"`
	Dimension  *string `json:"dimension,omitempty"`
}

// =======================
// ====    Methods    ====
// =======================
func (b *Booking) GetWeekday() time.Weekday {
	return (b.Start.Weekday() + 6) % 7
}

func ExtractUserIdsDuplicates(schedulerData *SchedulerData) []string {
	var results []string
	for _, user := range schedulerData.Users {
		for i := 0; i < *user.OfficeDays; i++ {
			results = append(results, *user.Id)
		}
	}
	return results
}
