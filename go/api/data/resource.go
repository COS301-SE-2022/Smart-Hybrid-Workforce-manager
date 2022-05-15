package data

import (
	"api/db"
	"database/sql"
	"strings"
	// "database/sql"
)

//////////////////////////////////////////////////
// Structures and Variables

//ResourceType enum
type ResourceType int8

const (
	PARKING ResourceType = iota + 1
	DESK
	MEETINGROOM
)

// Building identifies a Building Resource via common attributes
type Building struct {
	Id        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	Location  *string `json:"location,omitempty"`
	Dimension *string `json:"dimension,omitempty"`
}

// Buildings represent a splice of Building
type Buildings []*Building

// Room identifies a Room Resource via common attributes
type Room struct {
	Id         *string `json:"id,omitempty"`
	BuildingId *string `json:"building_id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Location   *string `json:"location,omitempty"`
	Dimension  *string `json:"dimension,omitempty"`
}

// Rooms represent a splice of Room
type Rooms []*Room

// RoomAssociation identifies a Room Association Resource via common attributes
type RoomAssociation struct {
	RoomId            *string `json:"room_id,omitempty"`
	RoomIdAssociation *string `json:"room_id_association,omitempty"`
}

// RoomAssociations represent a splice of RoomAssociation
type RoomAssociations []*RoomAssociation

// Resource identifies a Resource via common attributes
type Resource struct {
	Id           *string       `json:"id,omitempty"`
	RoomId       *string       `json:"room_id,omitempty"`
	Name         *string       `json:"name,omitempty"`
	Location     *string       `json:"location,omitempty"`
	RoleId       *string       `json:"role_id,omitempty"`
	ResourceType *ResourceType `json:"resource_type,omitempty"`
	DateCreated  *string       `json:"date_created,omitempty"`
}

// Resources represent a splice of Resource
type Resources []*Resource

// BookingDA provides access to the database for authentication purposes
type ResourceDA struct {
	access *db.Access
}

// NewBookingDA creates a new data access from a underlying shared data access
func NewResourceDA(access *db.Access) *ResourceDA {
	return &ResourceDA{
		access: access,
	}
}

//////////////////////////////////////////////////
// Mappers

func mapBuilding(rows *sql.Rows) (interface{}, error) {
	var identifier Building
	err := rows.Scan(
		&identifier.Id,
		&identifier.Name,
		&identifier.Location,
		&identifier.Dimension,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

func mapRoom(rows *sql.Rows) (interface{}, error) {
	var identifier Room
	err := rows.Scan(
		&identifier.Id,
		&identifier.BuildingId,
		&identifier.Name,
		&identifier.Location,
		&identifier.Dimension,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

func mapRoomAssociation(rows *sql.Rows) (interface{}, error) {
	var identifier RoomAssociation
	err := rows.Scan(
		&identifier.RoomId,
		&identifier.RoomIdAssociation,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

func ParseStringToResourceType(str string) ResourceType {
	resourceTypeMap := map[string]ResourceType{ // TODO [KP]: Move this out of the function
		"DESK":        DESK,
		"PARKING":     PARKING,
		"MEETINGROOM": MEETINGROOM,
	}
	r, _ := resourceTypeMap[strings.ToLower(str)]
	return r
}

func mapResource(rows *sql.Rows) (interface{}, error) {
	var tmp *string
	var identifier Resource
	err := rows.Scan(
		&identifier.Id,
		&identifier.RoomId,
		&identifier.Name,
		&identifier.Location,
		&identifier.RoleId,
		&tmp,
		&identifier.DateCreated,
	)
	if err != nil {
		return nil, err
	}

	temp := ParseStringToResourceType(*tmp)
	identifier.ResourceType = &temp

	return identifier, nil
}

//////////////////////////////////////////////////
// Functions

////////////////
// Building

// StoreBuildingResource stores a building
func (access *ResourceDA) StoreBuildingResource(identifier *Building) error {
	_, err := access.access.Query(
		`SELECT 1 FROM resource.building_store($1, $2, $3, $4)`, nil,
		identifier.Id, identifier.Name, identifier.Location, identifier.Dimension)
	if err != nil {
		return err
	}
	return nil
}

//FindBuildingResource finds a building
func (access *ResourceDA) FindBuildingResource(identifier *Building) (Buildings, error) {
	results, err := access.access.Query(
		`SELECT * FROM resource.building_find($1, $2, $3, $4)`, mapBuilding,
		identifier.Id, identifier.Name, identifier.Location, identifier.Dimension)
	if err != nil {
		return nil, err
	}
	tmp := make([]*Building, 0)
	for r, _ := range results {
		if value, ok := results[r].(Building); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeleteBuildingResource finds a building
func (access *ResourceDA) DeleteBuildingResource(identifier *Building) (*Building, error) {
	results, err := access.access.Query(
		`SELECT * FROM resource.building_remove($1)`, mapBuilding,
		identifier.Id)
	if err != nil {
		return nil, err
	}
	var tmp Buildings
	tmp = make([]*Building, 0)
	for r, _ := range results {
		if value, ok := results[r].(Building); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Building
func (buildings Buildings) FindHead() *Building {
	if len(buildings) == 0 {
		return nil
	}
	return buildings[0]
}

////////////////
// Room

// StoreRoomResource stores a Room
func (access *ResourceDA) StoreRoomResource(identifier *Room) error {
	_, err := access.access.Query(
		`SELECT 1 FROM resource.room_store($1, $2, $3, $4, $5)`, nil,
		identifier.Id, identifier.BuildingId, identifier.Name, identifier.Location, identifier.Dimension)
	if err != nil {
		return err
	}
	return nil
}

//FindRoomResource finds a Room
func (access *ResourceDA) FindRoomResource(identifier *Room) (Rooms, error) {
	results, err := access.access.Query(
		`SELECT * FROM resource.room_find($1, $2, $3, $4, $5)`, mapRoom,
		identifier.Id, identifier.BuildingId, identifier.Name, identifier.Location, identifier.Dimension)
	if err != nil {
		return nil, err
	}
	tmp := make([]*Room, 0)
	for r, _ := range results {
		if value, ok := results[r].(Room); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeleteRoomResource finds a Room
func (access *ResourceDA) DeleteRoomResource(identifier *Room) (*Room, error) {
	results, err := access.access.Query(
		`SELECT * FROM resource.room_remove($1)`, mapRoom,
		identifier.Id)
	if err != nil {
		return nil, err
	}
	var tmp Rooms
	tmp = make([]*Room, 0)
	for r, _ := range results {
		if value, ok := results[r].(Room); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Room
func (Rooms Rooms) FindHead() *Room {
	if len(Rooms) == 0 {
		return nil
	}
	return Rooms[0]
}
