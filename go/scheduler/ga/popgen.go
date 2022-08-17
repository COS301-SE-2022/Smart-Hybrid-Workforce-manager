package ga

import (
	cu "lib/collectionutils"
	"lib/utils"
)

func StubPopulationGenerator(domain *Domain, popSize int) Individuals {
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

func DayVResourcePopulationGenerator(domain *Domain, popSize int) Individuals {
	// Using representation 1 (DayVResource)
	schedulerData := domain.SchedulerData

	const numDaysInWeek int = 7 // Assume 7 days in a week for now

	// Get slot sizes/day
	slotSizes := []int{}
	for i := 0; i < numDaysInWeek; i++ { // TODO: @JonathanEnslin find out if we want to include all 7 days of week
		slotSizes = append(slotSizes, len(domain.SchedulerData.Resources)) // Initially assume all slots would be open
	}
	for _, booking := range *schedulerData.CurrentBookings {
		day := int(booking.GetWeekday())
		slotSizes[day]-- // Minus one slot for each existing booking on a day
	}
	// Make collection of available slots (used to select rand open slots later)
	// 2D where each row is a day
	openSlots := make([][]int, 0)
	for dayOfWeek := 0; dayOfWeek < numDaysInWeek; dayOfWeek++ {
		openSlots = append(openSlots, []int{})
		for i := 0; i < slotSizes[dayOfWeek]; i++ {
			openSlots[dayOfWeek] = append(openSlots[dayOfWeek], i)
		}
	}

	// Keep track of how many days users are already coming into the office
	timesAlreadyComingIn := make(map[string]int, 0) // (map[user id]times coming in already)
	for _, booking := range *schedulerData.CurrentBookings {
		timesAlreadyComingIn[*booking.UserId]++ // Add one to indicate they are coming in
	}

	// Add users as many times as they need to come into office
	usersToAdd := []string{}
	daysAvailable := make(map[string][]int, 0) // Used to indicate what days they are not already coming in (map[user id][days available])
	for _, user := range domain.SchedulerData.Users {
		daysAvailable[*user.Id] = cu.SequentialSequence(0, 7)
		// Add them times they have to come in - days they already come in
		for i := 0; i < *user.OfficeDays-timesAlreadyComingIn[*user.Id]; i++ { // TODO: @JonathanEnslin find out what do if no office days?
			usersToAdd = append(usersToAdd, *user.Id)
		}
	}

	// Merge this into other loop for cleaner less verbose code!!!
	// Remove days that users already come in from their available days
	for _, booking := range *schedulerData.CurrentBookings {
		bUserId := *booking.UserId
		if _, ok := daysAvailable[bUserId]; ok {
			// remove
			daysAvailable[bUserId] = cu.RemoveElementNoOrder(daysAvailable[bUserId], int(booking.GetWeekday()))
		}
	}

	openDays := cu.SequentialSequence(0, 7) // Used to keep track of what days still have space open
	for i := range openSlots {
		if len(openSlots[i]) == 0 {
			openDays = cu.RemoveElementNoOrder(openDays, i)
		}
	}

	openDaysCopy := []int{}
	population := make([]*Individual, domain.Config.Size)
	for i := 0; i < domain.Config.Size; i++ {
		copy(openDaysCopy, openDays)
		// Used to maintain a copy of openSlots
		openSlotsCopy := make([][]int, len(openSlots))
		for k := range openSlotsCopy {
			copy(openSlotsCopy[k], openSlots[k])
		}
		individual := Individual{Gene: make([][]string, numDaysInWeek)}
		// Initialise individual to empty indiv
		for j := range individual.Gene {
			individual.Gene[j] = make([]string, slotSizes[j])
		}

		// Randomly assign employees to slots
		for _, userId := range usersToAdd {
			availableDaysIntersection := cu.IntSliceIntersection(daysAvailable[userId], openDaysCopy)
			randDay := availableDaysIntersection[utils.RandInt(0, len(availableDaysIntersection))]
			// Get random slot index
			randSloti := utils.RandInt(0, len(openSlotsCopy[randDay]))
			randSlot := openSlotsCopy[randDay][randSloti]

			individual.Gene[randDay][randSlot] = userId // Assign the slot

			// Update structures indicating available slots/days
			daysAvailable[userId] = cu.RemoveElementNoOrder(daysAvailable[userId], randDay)
			openSlotsCopy[randDay] = cu.RemElemenAtI(openSlotsCopy[randDay], randSloti)
			if len(openSlotsCopy[randDay]) == 0 { // If no more spots open
				openDaysCopy = cu.RemoveElementNoOrder(openDaysCopy, randDay)
			}
		}
		population[i] = &individual
	}
	return population
}
