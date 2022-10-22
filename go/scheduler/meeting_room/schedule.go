package meetingroom

import (
	cu "lib/collectionutils"
	"scheduler/data"
	"sort"
)

// Extracts bookings that still need to be scheduled
func extractUnscheduledMeetingRoomBookings(schedulerData *data.SchedulerData) data.Bookings {
	unscheduled := data.Bookings{}
	for _, booking := range *schedulerData.CurrentBookings {
		if booking.ResourceId == nil && *booking.ResourceType == "MEETINGROOM" {
			unscheduled = append(unscheduled, booking)
		}
	}
	return unscheduled
}

// Extracts bookings that are not dependant on other bookings, as well as maps these bookings ids to dependant bookings
func extractIndependantBookings(bookings data.Bookings) (independants data.Bookings, dependants map[string]data.Bookings) {
	independants = data.Bookings{}
	dependants = make(map[string]data.Bookings)
	for _, booking := range bookings {
		if booking.Dependent == nil || *booking.Dependent == *booking.Id {
			independants = append(independants, booking)
		} else {
			if !cu.MapHasKey(dependants, *booking.Dependent) {
				dependants[*booking.Dependent] = data.Bookings{}
			}
			dependants[*booking.Dependent] = append(dependants[*booking.Dependent], booking)
		}
	}
	return
}

// Extracts all the meeting room resources
func extractMeetingRooms(schedulerData *data.SchedulerData) data.Resources {
	meetingRooms := data.Resources{}
	for _, resource := range schedulerData.Resources {
		if *resource.ResourceType == "MEETINGROOM" {
			meetingRooms = append(meetingRooms, resource)
		}
	}
	return meetingRooms
}

func MapMeetingRoomBookings(schedulerData *data.SchedulerData) map[string]*data.MeetingRoomBooking {
	mappedBookings := make(map[string]*data.MeetingRoomBooking)
	for _, meetingRoomBooking := range schedulerData.MeetingRoomBookings {
		mappedBookings[*meetingRoomBooking.BookingId] = meetingRoomBooking
	}
	return mappedBookings
}

// Sorts the resources by capacity in ascending order
func SortResources(resources data.Resources) {
	sort.SliceStable(resources, func(i, j int) bool {
		return resources[i].GetCapacity() < resources[j].GetCapacity()
	})
}

func GetNeededCapacity(schedulerData *data.SchedulerData, booking *data.MeetingRoomBooking) int {
	capacityNeeded := 0
	if booking.AdditionalAttendees != nil {
		capacityNeeded += *booking.AdditionalAttendees
	}
	if booking.TeamId != nil {
		capacityNeeded += schedulerData.TeamSize(*booking.TeamId)
	}
	if booking.RoleId != nil {
		capacityNeeded += schedulerData.RoleSize(*booking.RoleId)
	}
	return capacityNeeded
}

func AssignResources(schedulerData *data.SchedulerData) (scheduledBookings data.Bookings) {
	scheduledBookings = data.Bookings{}
	unscheduledBookings := extractUnscheduledMeetingRoomBookings(schedulerData)
	independants, dependants := extractIndependantBookings(unscheduledBookings)
	mappedBookings := MapMeetingRoomBookings(schedulerData)
	meetingRooms := extractMeetingRooms(schedulerData)
	// Sort the resources
	SortResources(meetingRooms)

	trueVar := true

	// For each independant booking, assign an available
	for _, independant := range independants {
		mappedBooking := mappedBookings[*independant.Id]
		capacityNeeded := GetNeededCapacity(schedulerData, mappedBooking)
		// Find a large enough resource
		remIndex := -1
		for i, room := range meetingRooms {
			if room.GetCapacity() >= capacityNeeded {
				// Assign the room
				independant.ResourceId = room.Id
				independant.Booked = &trueVar
				scheduledBookings = append(scheduledBookings, independant)
				remIndex = i
				break
			}
		}
		if remIndex == -1 { // If no resource was assigned
			// Assign a resource without a capacity
			for i, room := range meetingRooms {
				if room.GetCapacity() == -1 {
					// Assign the room
					independant.ResourceId = room.Id
					independant.Booked = &trueVar
					scheduledBookings = append(scheduledBookings, independant)
					remIndex = i
					break
				}
			}
		}
		if remIndex != -1 {
			// Remove this resource, while preserving order
			meetingRooms = cu.RemElementAtIPreseveOrder(meetingRooms, remIndex)
		}
	}
	scheduledIndependants := data.Bookings{}
	// Assign resource to all dependant bookings
	for _, scheduledIndependant := range scheduledBookings {
		for _, dependant := range dependants[*scheduledIndependant.Id] {
			dependant.ResourceId = scheduledIndependant.ResourceId
			dependant.Booked = &trueVar
			scheduledIndependants = append(scheduledIndependants, dependant)
		}
	}

	// Returns all assigned resources
	return append(scheduledBookings, scheduledIndependants...)
}

func AssignMeetingRoomsToBookings(schedulerData *data.SchedulerData) []data.Bookings {
	return []data.Bookings{AssignResources(schedulerData)}
}
