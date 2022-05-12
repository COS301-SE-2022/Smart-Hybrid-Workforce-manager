package data

import (
	"api/db"
	// "database/sql"
)

//////////////////////////////////////////////////
// Structures and Variables
type ResourceType int8

const (
	Parking ResourceType = iota + 1
	Desk
	MeetingRoom
)

// ResourceType func to get as string
func (r ResourceType) String() string {
	return [...]string{"PARKING", "DESK", "MEETINGROOM"}[r-1]
}

// Building indentifies a Building Resource via common attributes
type Building struct {
	Id        *string `json:"id,omitempty"`
	Location  *string `json:"location,omitempty"`
	Dimension *string `json:"dimension,omitempty"`
}

type Room struct {
	Id             *string  `json:"id,omitempty"`
	BuildingId     *string  `json:"building_id,omitempty"`
	Location       *string  `json:"location,omitempty"`
	Dimension      *string  `json:"dimension,omitempty"`
	RoomAssociates []string `json:"room_associates,omitempty"`
}

// BookingDA provides access to the database for authentication purposes
type ResourcesDA struct {
	access *db.Access
}

// NewBookingDA creates a new data access from a underlying shared data access
func NewResourceDA(access *db.Access) *ResourcesDA {
	return &ResourcesDA{
		access: access,
	}
}

//////////////////////////////////////////////////
// Mappers
// func mapResource(rows *sql.Rows) (interface{}, error){
	
// }


//////////////////////////////////////////////////
// Functions
