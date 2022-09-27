package data

import (
	"time"
)

func TimeInIntervalInclusive(check time.Time, start time.Time, end time.Time) bool {
	return (check.After(start) || start == check) && (check.Before(end) || end == check)
}

// Returns true if the resource is not allocated in a certain interval
func (schedulerData *SchedulerData) ResourceAvailable(resource *Resource, from time.Time, to time.Time) bool {
	for _, booking := range *schedulerData.CurrentBookings {
		if booking.ResourceId == resource.Id {
			if TimeInIntervalInclusive(*booking.Start, from, to) || TimeInIntervalInclusive(*booking.End, from, to) {
				return false
			}
		}
	}
	return true
}

func (schedulerData *SchedulerData) TeamSize(teamId string) int {
	return len(schedulerData.TeamsMap[teamId].UserIds)
}

func (schedulerData *SchedulerData) RoleSize(RoleId string) int {
	return len(schedulerData.RolesMap[RoleId].UserIds)
}
