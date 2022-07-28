package data

import (
	"api/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

//////////////////////////////////////////////////
// Structures and Variables

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

func fancyStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (b *MeetingRoomBooking) String() string {
	bookingStr := fmt.Sprint("Booking: ", b.Booking)
	idStr := "Id: " + fancyStr(b.Id)
	bookingIdStr := "BookingId: " + fancyStr(b.BookingId)
	teamIdStr := "TeamId: " + fancyStr(b.TeamId)
	return "\n========================\n" + bookingStr + "\n" + idStr + "\n" + bookingIdStr + "\n" + teamIdStr + "\n========================\n"
}

// MeetingRoomBookings represents a splice of BookingMeetingRooms
type MeetingRoomBookings []*MeetingRoomBooking

// Bookings represent a splice of Booking
type Bookings []*Booking

// BookingDA provides access to the database for authentication purposes
type BookingDA struct {
	access *db.Access
}

// NewBookingDA creates a new data access from a underlying shared data access
func NewBookingDA(access *db.Access) *BookingDA {
	return &BookingDA{
		access: access,
	}
}

// Commit commits the current implicit transaction
func (access *BookingDA) Commit() error {
	return access.access.Commit()
}

//////////////////////////////////////////////////
// Mappers

func mapIdReturn(rows *sql.Rows) (interface{}, error) {
	var id string
	err := rows.Scan(&id)
	if err != nil {
		return nil, err
	}
	return id, err
}

func mapBooking(rows *sql.Rows) (interface{}, error) {
	var identifier Booking
	err := rows.Scan(
		&identifier.Id,
		&identifier.UserId,
		&identifier.ResourceType,
		&identifier.ResourcePreferenceId,
		&identifier.ResourceId,
		&identifier.Start,
		&identifier.End,
		&identifier.Booked,
		&identifier.DateCreated,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

func mapMeetingRoomBooking(rows *sql.Rows) (interface{}, error) {
	var meetingRoomBooking MeetingRoomBooking
	err := rows.Scan(
		&meetingRoomBooking.Id,
		&meetingRoomBooking.BookingId,
		&meetingRoomBooking.TeamId,
		&meetingRoomBooking.RoleId,
		&meetingRoomBooking.AdditionalAttendees,
		&meetingRoomBooking.DesksAttendees,
		&meetingRoomBooking.DesksAdditionalAttendees,
	)
	if err != nil {
		return nil, err
	}
	return meetingRoomBooking, nil
}

//////////////////////////////////////////////////
// Functions

// StoreIdentifier stores an identifier
func (access *BookingDA) StoreIdentifier(identifier *Booking) (*string, error) {
	id, err := access.access.Query(
		`SELECT * FROM booking.identifier_store($1, $2, $3, $4, $5, $6, $7, $8)`, mapIdReturn,
		identifier.Id, identifier.UserId, identifier.ResourceType, identifier.ResourcePreferenceId, identifier.ResourceId, identifier.Start, identifier.End, identifier.Booked)
	if err != nil {
		return nil, err
	}
	_id := fmt.Sprint(id[0])
	return &_id, nil
}

//FindIdentifier finds an identifier
func (access *BookingDA) FindIdentifier(identifier *Booking, permissions *Permissions) (Bookings, error) {
	permissionContent, err := json.Marshal(*permissions)
	if err != nil {
		return nil, err
	}
	results, err := access.access.Query(
		`SELECT * FROM booking.identifier_find($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, mapBooking,
		identifier.Id, identifier.UserId, identifier.ResourceType, identifier.ResourcePreferenceId,
		identifier.ResourceId, identifier.Start, identifier.End, identifier.Booked, identifier.DateCreated, permissionContent)
	if err != nil {
		return nil, err
	}
	tmp := make([]*Booking, 0)
	for r := range results {
		if value, ok := results[r].(Booking); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeleteIdentifier finds an identifier
func (access *BookingDA) DeleteIdentifier(identifier *Booking) (*Booking, error) {
	results, err := access.access.Query(
		`SELECT * FROM booking.identifier_remove($1)`, mapBooking,
		identifier.Id)
	if err != nil {
		return nil, err
	}
	var tmp Bookings
	tmp = make([]*Booking, 0)
	for r := range results {
		if value, ok := results[r].(Booking); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Booking
func (bookings Bookings) FindHead() *Booking {
	if len(bookings) == 0 {
		return nil
	}
	return bookings[0]
}

// StoreBookingMeetingRoom stores or updates a meeting room booking
func (access *BookingDA) StoreBookingMeetingRoom(meetingRoomBooking *MeetingRoomBooking) error {
	_, err := access.access.Query(
		`SELECT 1 FROM booking.meeting_room_store($1, $2, $3, $4, $5, $6, $7)`, nil,
		meetingRoomBooking.Id, meetingRoomBooking.BookingId, meetingRoomBooking.TeamId, meetingRoomBooking.RoleId, meetingRoomBooking.AdditionalAttendees, meetingRoomBooking.DesksAttendees, meetingRoomBooking.DesksAdditionalAttendees)
	if err != nil {
		return err
	}
	return nil
}

// FindMeeetingRoomBooking finds an identifier
func (access *BookingDA) FindMeeetingRoomBooking(meetingRoomBooking *MeetingRoomBooking, permissions *Permissions) (MeetingRoomBookings, error) {
	permissionContent, err := json.Marshal(*permissions)
	if err != nil {
		return nil, err
	}
	results, err := access.access.Query(
		`SELECT * FROM booking.meeting_room_find($1, $2, $3, $4, $5, $6, $7, $8)`, mapMeetingRoomBooking,
		meetingRoomBooking.Id, meetingRoomBooking.BookingId, meetingRoomBooking.TeamId, meetingRoomBooking.RoleId,
		meetingRoomBooking.AdditionalAttendees, meetingRoomBooking.DesksAttendees, meetingRoomBooking.DesksAdditionalAttendees, permissionContent)
	if err != nil {
		return nil, err
	}
	tmp := make([]*MeetingRoomBooking, 0)
	for r := range results {
		if value, ok := results[r].(MeetingRoomBooking); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeleteIdentifier finds an identifier
func (access *BookingDA) DeleteMeetingRoomBooking(meetingRoomBooking *Booking) (*MeetingRoomBooking, error) {
	results, err := access.access.Query(
		`SELECT * FROM booking.meeting_room_remove($1)`, mapMeetingRoomBooking,
		meetingRoomBooking.Id)
	if err != nil {
		return nil, err
	}
	var tmp MeetingRoomBookings
	tmp = make([]*MeetingRoomBooking, 0)
	for r := range results {
		if value, ok := results[r].(MeetingRoomBooking); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Booking
func (meetingRoomBookings MeetingRoomBookings) FindHead() *MeetingRoomBooking {
	if len(meetingRoomBookings) == 0 {
		return nil
	}
	return meetingRoomBookings[0]
}
