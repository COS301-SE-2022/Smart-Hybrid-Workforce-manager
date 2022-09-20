package ga

import (
	cu "lib/collectionutils"
	"math"
	"scheduler/data"
)

///////////////////////////////////////////////////
// WEEKLY

func WeeklyStubFitness(domain *Domain, individuals Individuals) []float64 {
	var result []float64
	for i := 0; i < len(individuals); i++ {
		result = append(result, 0.0)
		individuals[i].Fitness = 0.0
	}
	return result
}

func WeeklyDayVResourceFitness(domain *Domain, individuals Individuals) []float64 {
	var result []float64
	for _, individual := range individuals {
		fitness := weeklyDayVResourceFitness(domain, individual)
		result = append(result, fitness)
		individual.Fitness = fitness
		//fmt.Println(fitness)
	}
	return result
}

func weeklyDayVResourceFitness(domain *Domain, individual *Individual) float64 {
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
	prevWeek := make(map[string]([]bool)) // map[userId]{array representing a week, true means user comes in on day}
	thisWeek := make(map[string]([]bool)) // map[userId]{array representing a week, true means user comes in on day}

	// Find what days users came in previously
	for _, booking := range *schedulerData.PastBookings {
		if len(prevWeek[*booking.UserId]) == 0 {
			prevWeek[*booking.UserId] = make([]bool, numDaysInWeek)
		}
		prevWeek[*booking.UserId][booking.GetWeekday()] = true
	}
	// Find which days they come in this week
	for dayi, day := range individual.Gene {
		for _, uid := range day {
			// fmt.Printf("... %d <<<<<< %d\n", numDaysInWeek, dayi)
			if len(thisWeek[uid]) == 0 {
				thisWeek[uid] = make([]bool, numDaysInWeek)
			}
			// fmt.Printf("%d <<<<<< %d  len(%d)\n", numDaysInWeek, dayi, len(thisWeek[uid]))
			thisWeek[uid][dayi] = true
		}
	}
	dissimilarityCount := 0
	// Count how many days don't overlap
	for key := range prevWeek {
		for i := 0; i < numDaysInWeek; i++ {
			// fmt.Println(len(prevWeek[key]), "   ", thisWeek[key])
			// Check if how many days the come in where they did not come in on the same day in the previous week
			if len(prevWeek[key]) == len(thisWeek[key]) && prevWeek[key][i] != thisWeek[key][i] {
				dissimilarityCount += 1
			}
			if len(prevWeek[key]) != len(thisWeek[key]) && len(prevWeek[key]) == 0 {
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
			if entry != "" {
				// //fmt.Println(freq[entry])
				freq[entry] = freq[entry] + 1
			}
		}
		result = append(result, freq)
	}
	return result
}

///////////////////////////////////////////////////
// DAILY

func DailyFitness(domain *Domain, individuals Individuals) []float64 {
	var result []float64
	for _, individual := range individuals {
		fitness := dailyFitness(domain, individual)
		result = append(result, fitness)
		individual.Fitness = fitness
		//fmt.Println(fitness)
	}
	return result
}

func dailyFitness(domain *Domain, individual *Individual) float64 {
	return 0.0
}

type teamRoomGroups struct {
	teamId string
	// A map mapping roomId to the members part of the team in that room (map[roomId]{arr of user indices in that room})
	roomGroups map[string][]int
}

type teamRoomProximity struct {
	teamRoomGroups
	// A map mapping roomIds to proximity scores for members in that room
	roomProximities map[string]float64
}

// preferredDeskBonus returns a bonus fitness value for users sitting at their preffered desk
func (individual *Individual) preferredDeskBonuses(domain *Domain) float64 {
	gene := individual.Gene
	bonusSum := 0.0
	for i := 0; i < len(gene[0]); i++ {
		// Get the users preferred resource
		userPrefResourceProx := userPreferredDeskProximity(individual, domain, i)
		if userPrefResourceProx >= 0 {
			userPrefResourceProx += 1.0 // In case it is 0
			bonusSum += 1 / userPrefResourceProx
		}
	}
	return bonusSum / float64(len(gene[0])) // Take the average score, this should usually keep the value below the team scores
}

// Calculates the distance the user is from their preferred desk
// returns -1.0 if the user does not have a preferred resource, -2.0 if the user
// is in a different room than the preferred resource
func userPreferredDeskProximity(indiv *Individual, domain *Domain, userIndex int) float64 {
	// Get the users coordinates and roomId
	userCoords := indiv.getUserCoordinate(domain, userIndex)
	userRoomId := domain.SchedulerData.ResourcesMap[indiv.Gene[0][userIndex]].RoomId
	// Get the users preferred resources
	preferredResource := getUserPreferredResource(domain, userIndex)
	if preferredResource == nil { // If the user has no preferred resource
		return -1.0
	}
	if *preferredResource.RoomId != *userRoomId {
		return -2.0
	}
	resourceCoords := []float64{*preferredResource.XCoord, *preferredResource.YCoord}
	return math.Sqrt(distanceRadicand(userCoords, resourceCoords))
}

// Gets the users preferred resource, or nil if it does not exist
func getUserPreferredResource(domain *Domain, userIndex int) *data.Resource {
	userId := domain.Map[userIndex]              // Get the users Id
	user := domain.SchedulerData.UserMap[userId] // Get the user
	if user.PreferredDesk == nil {               // If the user has not preferred desk
		return nil
	}
	if !cu.MapHasKey(domain.SchedulerData.ResourcesMap, *user.PreferredDesk) { // If the preferred desk does not exist
		return nil
	}
	return domain.SchedulerData.ResourcesMap[*user.PreferredDesk]
}

// Gets the priority of the team, returns -1 if priority is nil
func getTeamPriority(domain *Domain, teamId string) int {
	prio := domain.SchedulerData.TeamsMap[teamId].Priority
	if prio == nil {
		return -1
	}
	return *prio
}

// teamProximityScore calculates a score that indicates the proximity of members
// of a team, scales with team priority (TODO: @JonathanEnslin remember this)
func (individual *Individual) teamProximityScore(domain *Domain) float64 {
	teamRoomProximities := individual.getTeamRoomProximities(domain)
	scores := make([]float64, len(teamRoomProximities))
	// TODO: @JonathanEnslin filter out empties and stuff if necessary
	for i, teamRoomProx := range teamRoomProximities {
		// Use reciprocal, since if the teams have a larger avg distance from the centroid
		// the fitness should be smaller
		scores[i] = math.Max(1.0, float64(getTeamPriority(domain, teamRoomProx.teamId))) / (individualTeamProximityScore(teamRoomProx) + 1.0)
	}
	// Sum all the reciprocals
	return cu.Sum(scores)
}

func individualTeamProximityScore(teamRoomProx teamRoomProximity) float64 {
	// Sum over the proximity of all rooms
	sum := 0.0
	for _, prox := range teamRoomProx.roomProximities {
		sum += prox
	}
	// TODO: @JonathanEnslin add penalty for teams split over rooms
	return sum
}

// Takes in a index of user, and returns the coordinates of the user according to the resource
// that is assigned to them
func (individual *Individual) getUserCoordinate(domain *Domain, index int) []float64 {
	resource := domain.SchedulerData.ResourcesMap[individual.Gene[0][index]]
	return []float64{
		*resource.XCoord,
		*resource.YCoord,
	} // (x, y)
}

// Calculates the proximities of teams grouped by the rooms they are in
func (individual *Individual) getTeamRoomProximities(domain *Domain) []teamRoomProximity {
	// Get team room groups
	teamRoomGroupsArr := individual.getTeamsGroupedByRooms(domain)
	// Create empty teamRoomProximitySlice
	teamRoomProximities := make([]teamRoomProximity, len(teamRoomGroupsArr))
	// Function used to compile a slice of coordinates from user indices
	compileCoordinates := func(userIndices []int) [][]float64 {
		// Allocate slice
		coords := make([][]float64, len(userIndices))
		for i, index := range userIndices {
			// Get the coordinates for each user
			coords[i] = individual.getUserCoordinate(domain, index)
		}
		return coords
	}
	for i, teamRoomGroup := range teamRoomGroupsArr {
		// Allocate map
		roomProximites := make(map[string]float64, len(teamRoomGroup.roomGroups))
		for roomId, usersInRooms := range teamRoomGroup.roomGroups {
			// For each room, get the proximity by compiling the coordinates and getting avg dist from centroid
			roomProximites[roomId] = avgDistanceFromCentroid(compileCoordinates(usersInRooms))
		}
		teamRoomProximities[i] = teamRoomProximity{
			teamRoomGroups:  teamRoomGroup,
			roomProximities: roomProximites,
		}
	}
	return teamRoomProximities
}

// getTeamsGroupedByRooms returns
func (individual *Individual) getTeamsGroupedByRooms(domain *Domain) []teamRoomGroups {
	schedulerData := domain.SchedulerData
	gene := individual.Gene
	// =================================================================
	// TODO: @JonathanEnslin look at moving this piece of code into the domain as well since it is common across individuals
	// A map that contains the user indices per team
	teamUserIndices := domain.GetTeamUserIndices()
	// =================================================================

	// Group each team by room
	roomGroupingFunc := func(userIndex int) string { // returns a roomId
		return *schedulerData.ResourcesMap[gene[0][userIndex]].RoomId // Return the room that the user will be in
	}
	groups := []teamRoomGroups{}
	for teamId, indices := range teamUserIndices {
		_, roomGroups := cu.GroupBy(indices, roomGroupingFunc) // Get the room groups for the team
		groups = append(groups, teamRoomGroups{teamId: teamId, roomGroups: roomGroups})
	}
	return groups
}

// ================ Functions used for daily fitness ================

// avgDistanceFromCentroid calculates the centroid of a set of points and then calculates
// the avg point-to-centroid distance
// param coords is an array of float64 arrays, where each inner array corresponds to a set of coordinates
// the length of the array correspons to the coordinate dimension, e.g. len=2 would be 2D, or x and y
func avgDistanceFromCentroid(coords [][]float64) float64 {
	centroid := getCentroid(coords) // Get the centroid
	avgDistance := 0.0
	for _, coord := range coords {
		avgDistance += math.Sqrt(distanceRadicand(centroid, coord)) // get the total distance
	}
	avgDistance /= float64(len(coords)) // calculate the avg distance from the total
	return avgDistance
}

// Gets the centroid a set of coordinates
// See avgDistanceFromCentroid explanation on how the coordinates work
func getCentroid(coords [][]float64) []float64 {
	means := make([]float64, len(coords[0]))
	for _, coord := range coords {
		for i := range coord {
			means[i] += coord[i] // sum the total value for each part of the coordinate
		}
	}
	for i := range means {
		means[i] /= float64(len(coords)) // calculate the average from the totals
	}
	return means
}

// Returns the value of the distance formula before sqrt is applied
// See avgDistanceFromCentroid explanation on how the coordinates work
func distanceRadicand(origin []float64, coord []float64) float64 {
	result := 0.0
	for i := range coord {
		result += math.Pow(coord[i]-origin[i], 2)
	}
	return result
}
