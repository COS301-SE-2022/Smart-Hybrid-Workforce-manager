package ga

import (
	cu "lib/collectionutils"
	"lib/logger"
	"lib/testutils"
	"lib/utils"
)

///////////////////////////////////////////////////
// WEEKLY

func WeeklyStubPopulationGenerator(domain *Domain, popSize int) Individuals {
	var population Individuals
	for j := 0; j < popSize; j++ {
		genes := []string{}
		for i := 0; i < len(domain.SchedulerData.Users); i++ {
			genes = append(genes, domain.Terminals[j%len(domain.Terminals)])
		}
		population = append(population, &Individual{Gene: [][]string{genes}})
	}
	return population
}

func WeeklyDayVResourcePopulationGenerator(domain *Domain, popSize int) Individuals {
	// Using representation 1 (DayVResource)
	const numDaysInWeek int = 5 // Assume 5 days in a week for now

	// Get slot info
	slotSizes, openSlots, openDays := getSlotInfo(domain, numDaysInWeek)

	// Get user related info
	daysAvailable, usersToAdd := getUserAvailabilityInfo(domain, numDaysInWeek)

	openDaysCopy := make([]int, len(openDays))
	population := make([]*Individual, popSize)
	daysAvailableCopy := make(map[string][]int)
	for i := 0; i < popSize; i++ {
		// make copy of days available map
		for key, val := range daysAvailable {
			daysAvailableCopy[key] = make([]int, len(val))
			copy(daysAvailableCopy[key], val)
		}
		copy(openDaysCopy, openDays)

		// Used to maintain a copy of openSlots
		openSlotsCopy := make([][]int, len(openSlots))
		for k := range openSlotsCopy {
			openSlotsCopy[k] = make([]int, len(openSlots[k]))
			copy(openSlotsCopy[k], openSlots[k])
		}

		individual := Individual{Gene: make([][]string, numDaysInWeek)}
		// Initialise individual to empty indiv
		for j := range individual.Gene { // DEBUG how long is GENE?
			individual.Gene[j] = make([]string, slotSizes[j])
		}

		// Randomly assign employees to slots
		for _, userId := range usersToAdd {
			availableDaysIntersection := cu.IntSliceIntersection(daysAvailableCopy[userId], openDaysCopy)
			// If user can not come in, either because they already come in each day they can, or
			// if there are no resources left
			if len(availableDaysIntersection) == 0 {
				break
			}

			randDay := availableDaysIntersection[utils.RandInt(0, len(availableDaysIntersection))]

			// Get random slot index
			randSloti := utils.RandInt(0, len(openSlotsCopy[randDay]))
			randSlot := openSlotsCopy[randDay][randSloti]

			individual.Gene[randDay][randSlot] = userId // Assign the slot

			// Update structures indicating available slots/days
			daysAvailableCopy[userId] = cu.RemoveElementNoOrder(daysAvailableCopy[userId], randDay)
			openSlotsCopy[randDay] = cu.RemElemenAtI(openSlotsCopy[randDay], randSloti)
			if len(openSlotsCopy[randDay]) == 0 { // If no more spots open
				openDaysCopy = cu.RemoveElementNoOrder(openDaysCopy, randDay)
			}
		}
		population[i] = &individual
	}

	// DEBUG ================
	logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	logger.Error.Println(testutils.Scolour(testutils.GREEN, "Done"))
	logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// END DEBUG ================

	return population
}

// getSlotInfo used the data in the scheduler and determines the amount of slots available on each day, and the
// days with open slots, it returns the slotSizes per day, a 2D array where the index corresponds to the day,
// and each row contains a sequence of numbers, corresponding to an open slot, and openDays, an array indicating
// which days still have space open
func getSlotInfo(domain *Domain, numDaysInWeek int) (slotSizes []int, openSlots [][]int, openDays []int) {
	schedulerData := domain.SchedulerData
	// Get slot sizes/day
	slotSizes = []int{}
	for i := 0; i < numDaysInWeek; i++ { // TODO: @JonathanEnslin find out if we want to include all 7 days of week
		slotSizes = append(slotSizes, len(domain.SchedulerData.Resources)) // Initially assume all slots would be open
	}

	for _, booking := range *schedulerData.CurrentBookings {
		day := int(booking.GetWeekday())
		slotSizes[day]-- // Minus one slot for each existing booking on a day
	}

	// Make collection of available slots (used to select rand open slots later)
	// 2D where each row is a day
	openSlots = make([][]int, 0)
	for dayOfWeek := 0; dayOfWeek < numDaysInWeek; dayOfWeek++ {
		openSlots = append(openSlots, []int{})
		for i := 0; i < slotSizes[dayOfWeek]; i++ {
			openSlots[dayOfWeek] = append(openSlots[dayOfWeek], i)
		}
	}

	openDays = cu.SequentialSequence(0, numDaysInWeek) // Used to keep track of what days still have space open
	for i := range openSlots {
		if len(openSlots[i]) == 0 {
			openDays = cu.RemoveElementNoOrder(openDays, i)
		}
	}

	return slotSizes, openSlots, openDays
}

// getUserAvailibilityInfo uses the schedulerData to determine which days users are available to come in, as well as the days
// that the users are available to come into the office, it returns daysAvailable, which is a mop where the keys are user ids and
// the values are int arrays, where each element in the day is an int that corresponds to a day of the week that the user is
// available on, it also returns usersToAdd, an array of user ids, where the user id is present in the array as many times
// as the user needs to come into the office in the week
func getUserAvailabilityInfo(domain *Domain, numDaysInWeek int) (daysAvailable map[string][]int, usersToAdd []string) {
	schedulerData := domain.SchedulerData
	// Keep track of how many days users are already coming into the office
	timesAlreadyComingIn := make(map[string]int, 0) // (map[user id]times coming in already)
	for _, booking := range *schedulerData.CurrentBookings {
		timesAlreadyComingIn[*booking.UserId]++ // Add one to indicate they are coming in
	}

	// Add users as many times as they need to come into office
	usersToAdd = []string{}
	daysAvailable = make(map[string][]int, 0) // Used to indicate what days they are not already coming in (map[user id][days available])
	for _, user := range domain.SchedulerData.Users {
		daysAvailable[*user.Id] = cu.SequentialSequence(0, numDaysInWeek)
		// Add them times they have to come in - days they already come in
		for i := 0; i < *user.OfficeDays-timesAlreadyComingIn[*user.Id]; i++ { // TODO: @JonathanEnslin find out what do if no office days?
			usersToAdd = append(usersToAdd, *user.Id)
		}
	}

	// Remove days that users already come in from their available days
	for _, booking := range *schedulerData.CurrentBookings {
		bUserId := *booking.UserId
		if _, ok := daysAvailable[bUserId]; ok {
			// remove
			daysAvailable[bUserId] = cu.RemoveElementNoOrder(daysAvailable[bUserId], int(booking.GetWeekday())) // inefficient
		}
	}

	return daysAvailable, usersToAdd
}

///////////////////////////////////////////////////
// DAILY

// DailyPopulationGenerator generates random possibly invalid individuals
func DailyPopulationGenerator(domain *Domain, popSize int) Individuals {
	var population Individuals
	for j := 0; j < popSize; j++ {
		var individual Individual
		individual.Gene = append(individual.Gene, domain.GetRandomTerminalArrays(len(domain.Map)))
		population = append(population, &individual)
	}
	return population
}
