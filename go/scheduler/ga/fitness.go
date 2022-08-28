package ga

import (
	"scheduler/data"
)

func StubFitness(domain *Domain, individuals Individuals) []float64 {
	var result []float64
	for i := 0; i < len(individuals); i++ {
		result = append(result, 0.0)
		individuals[i].Fitness = 0.0
	}
	return result
}

func DayVResourceFitness(domain *Domain, individuals Individuals) []float64 {
	var result []float64
	for _, individual := range individuals {
		fitness := dayVResourceFitness(domain, individual)
		result = append(result, fitness)
		individual.Fitness = fitness
		//fmt.Println(fitness)
	}
	return result
}

func dayVResourceFitness(domain *Domain, individual *Individual) float64 {
	dailyMaps := individual.getUserCountMapsPerDay()
	// fmt.Printf("daily maps: %v\n", dailyMaps)
	differentUsersCount := individual.DifferentUsersCount(domain)
	// fmt.Printf("diff user counts: %v\n", differentUsersCount)
	doubleBookings := individual.DoubleBookingsCount(domain, dailyMaps)
	// fmt.Printf("double bookings: %v\n", doubleBookings)
	usersNotCommingInTheirSpecifiedAmountCount := individual.UsersNotCommingInTheirSpecifiedAmountCount(domain, dailyMaps)
	// fmt.Printf("usersNotCommingInTheirSpecifiedAmountCountmaps: %v\n", usersNotCommingInTheirSpecifiedAmountCount)
	teamsAttendingSameDays := individual.TeamsAttendingSameDays(domain, dailyMaps)
	// fmt.Printf("teamsAttendingSameDays: %v\n", teamsAttendingSameDays)
	if doubleBookings == 0 {
		doubleBookings = 1
	}
	if usersNotCommingInTheirSpecifiedAmountCount == 0 {
		usersNotCommingInTheirSpecifiedAmountCount = 1
	}
	if differentUsersCount == 0 {
		differentUsersCount = 1
	}
	if teamsAttendingSameDays == 0 {
		teamsAttendingSameDays = 1
	}
	return float64(differentUsersCount) * teamsAttendingSameDays / (float64(doubleBookings) * float64(usersNotCommingInTheirSpecifiedAmountCount))
}

// DifferentUsersCount takes the sum of the amount of times users come in on a different day as the week before
func (individual *Individual) DifferentUsersCount(domain *Domain) int {
	numDaysInWeek := len(individual.Gene)
	// fmt.Println("+++++++++++++++++++ NUM DAYS IN WEEK: ", numDaysInWeek)
	schedulerData := domain.SchedulerData
	// Create map containg days a user comes in
	comesInPrevWeek := make(map[string]([]bool)) // map[userId]{array representing a week, true means user comes in on day}
	comesInThisWeek := make(map[string]([]bool)) // map[userId]{array representing a week, true means user comes in on day}

	// Find what days users came in previously
	for _, booking := range *schedulerData.PastBookings {
		// if len(comesInPrevWeek[*booking.UserId]) == 0 {
		// 	comesInPrevWeek[*booking.UserId] = make([]bool, numDaysInWeek)
		// }
		comesInPrevWeek[*booking.UserId][booking.GetWeekday()] = true
	}
	// Find which days they come in this week
	for dayi, day := range individual.Gene {
		for _, uid := range day {
			// fmt.Printf("... %d <<<<<< %d\n", numDaysInWeek, dayi)
			if len(comesInThisWeek[uid]) == 0 {
				comesInThisWeek[uid] = make([]bool, numDaysInWeek)
			}
			// fmt.Printf("%d <<<<<< %d  len(%d)\n", numDaysInWeek, dayi, len(comesInThisWeek[uid]))
			comesInThisWeek[uid][dayi] = true
		}
	}
	dissimilarityCount := 0
	// Count how many days don't overlap
	for key := range comesInPrevWeek {
		for i := 0; i < numDaysInWeek; i++ {
			if comesInPrevWeek[key][i] != comesInThisWeek[key][i] {
				dissimilarityCount += 1
			}
		}
	}
	return dissimilarityCount
}

// TeamsAttendingSameDays loops over all teams and tallys the following:
// team_priority x ([%users that come in on same days] / [number of different days])
func (individual *Individual) TeamsAttendingSameDays(domain *Domain, dailyMaps []map[string]int) float64 {
	total := 0.0
	for _, team := range domain.SchedulerData.Teams {
		total += float64(*team.Priority) * percentageUsersOnSameDayInTeam(team, dailyMaps)
	}
	return total
}

//percentageUsersOnSameDayInTeam calculates the amount of users as a percentage that come in together
func percentageUsersOnSameDayInTeam(team *data.TeamInfo, dailyMaps []map[string]int) float64 {
	if len(team.UserIds) == 0 {
		return 0.0
	}
	usersNotTogether := 0
	for _, day := range dailyMaps {
		usersTogether := 0
		for _, teamMember := range team.UserIds {
			if day[teamMember] > 0 {
				usersTogether++
			}
		}
		if usersTogether == 0 {
			usersNotTogether++
		}
	}
	return (float64(len(team.UserIds)) - float64(usersNotTogether)) / float64(len(team.UserIds))
}

// DoubleBookingsCount counts the amount of users double booked for the entire week
func (individual *Individual) DoubleBookingsCount(domain *Domain, dailyMaps []map[string]int) int {
	count := 0
	for _, day := range dailyMaps {
		for _, value := range day {
			if value >= 2 {
				count++
			}
		}
	}
	return count
}

// UsersNotCommingInTheirSpecifiedAmountCount Counts the amount of users that do not come in their specified amounts
func (individual *Individual) UsersNotCommingInTheirSpecifiedAmountCount(domain *Domain, dailyMaps []map[string]int) int {
	count := 0
	for _, user := range domain.SchedulerData.Users {
		total := 0
		for _, day := range dailyMaps {
			total += day[*user.Id]
		}
		if total != *user.OfficeDays {
			count++
		}
	}
	return count
}

// GetUserCountMapsPerDay returns the frequencies that users attend each day at
func (individual *Individual) getUserCountMapsPerDay() []map[string]int {
	var result []map[string]int
	for _, day := range individual.Gene {
		freq := make(map[string]int)
		for _, entry := range day {
			freq[entry] = freq[entry] + 1
		}
		result = append(result, freq)
	}
	return result
}
