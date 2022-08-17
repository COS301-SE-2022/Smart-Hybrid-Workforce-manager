package ga

import (
	cu "lib/collectionutils"
	"lib/logger"
	"lib/testutils"
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

	// DEBUG ================
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "1-%v", slotSizes))
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// time.Sleep(100 * time.Millisecond)
	// END DEBUG ================

	for _, booking := range *schedulerData.CurrentBookings {
		logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "1.5-%v", booking))
		day := int(booking.GetWeekday())
		slotSizes[day]-- // Minus one slot for each existing booking on a day
	}

	// DEBUG ================
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "2-%v", slotSizes))
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// time.Sleep(100 * time.Millisecond)
	// END DEBUG ================

	// Make collection of available slots (used to select rand open slots later)
	// 2D where each row is a day
	openSlots := make([][]int, 0)
	for dayOfWeek := 0; dayOfWeek < numDaysInWeek; dayOfWeek++ {
		openSlots = append(openSlots, []int{})
		for i := 0; i < slotSizes[dayOfWeek]; i++ {
			openSlots[dayOfWeek] = append(openSlots[dayOfWeek], i)
		}
	}

	// DEBUG ================
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "3-%v", openSlots))
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// time.Sleep(100 * time.Millisecond)
	// END DEBUG ================

	// Keep track of how many days users are already coming into the office
	timesAlreadyComingIn := make(map[string]int, 0) // (map[user id]times coming in already)
	for _, booking := range *schedulerData.CurrentBookings {
		timesAlreadyComingIn[*booking.UserId]++ // Add one to indicate they are coming in
	}

	// DEBUG ================
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// b, err := json.MarshalIndent(timesAlreadyComingIn, "", "  ")
	// if err != nil {
	// 	logger.Error.Println(testutils.Scolourf(testutils.YELLOW, "4-%v", err))
	// }
	// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "4-----\n%s", string(b)))
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// time.Sleep(100 * time.Millisecond)
	// END DEBUG ================

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

	// DEBUG ================
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "5.1-%v", usersToAdd))
	// b, err = json.MarshalIndent(daysAvailable, "", "  ")
	// if err != nil {
	// 	logger.Error.Println(testutils.Scolourf(testutils.YELLOW, "5.2-%v", err))
	// }
	// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "5.2------\n%s", string(b)))
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// time.Sleep(100 * time.Millisecond)
	// END DEBUG ================

	// Merge this into other loop for cleaner less verbose code!!!
	// Remove days that users already come in from their available days
	for _, booking := range *schedulerData.CurrentBookings {
		bUserId := *booking.UserId
		if _, ok := daysAvailable[bUserId]; ok {
			// remove
			daysAvailable[bUserId] = cu.RemoveElementNoOrder(daysAvailable[bUserId], int(booking.GetWeekday())) // inefficient
		}
	}

	// DEBUG ================
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// b, err = json.MarshalIndent(daysAvailable, "", "  ")
	// if err != nil {
	// 	logger.Error.Println(testutils.Scolourf(testutils.YELLOW, "6-%v", err))
	// }
	// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "6------\n%s", string(b)))
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// time.Sleep(100 * time.Millisecond)
	// END DEBUG ================

	openDays := cu.SequentialSequence(0, 7) // Used to keep track of what days still have space open
	for i := range openSlots {
		if len(openSlots[i]) == 0 {
			openDays = cu.RemoveElementNoOrder(openDays, i)
		}
	}

	// DEBUG ================
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "7-%v", openDays))
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// time.Sleep(100 * time.Millisecond)
	// END DEBUG ================

	// DEBUG ================
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "12-%v", openDays))
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// END DEBUG ================
	openDaysCopy := make([]int, len(openDays))
	population := make([]*Individual, domain.Config.PopulationSize)
	daysAvailableCopy := make(map[string][]int)
	for i := 0; i < popSize; i++ {
		// make copy of days available map
		for key, val := range daysAvailable {
			daysAvailableCopy[key] = make([]int, len(val))
			copy(daysAvailableCopy[key], val)
		}
		copy(openDaysCopy, openDays)
		// DEBUG ================
		// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
		// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "11-%v", openDaysCopy))
		// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
		// END DEBUG ================
		// Used to maintain a copy of openSlots
		openSlotsCopy := make([][]int, len(openSlots))
		for k := range openSlotsCopy {
			openSlotsCopy[k] = make([]int, len(openSlots[k]))
			copy(openSlotsCopy[k], openSlots[k])
		}
		individual := Individual{Gene: make([][]string, numDaysInWeek)}
		// Initialise individual to empty indiv
		for j := range individual.Gene {
			individual.Gene[j] = make([]string, slotSizes[j])
		}

		// Randomly assign employees to slots
		for _, userId := range usersToAdd {
			// DEBUG ================
			// logger.Error.Println(testutils.Scolourf(testutils.BLUE, ">>>>>>>>>>>>>>%v<<<<<<<<<<<<<", loop))
			// time.Sleep(25 * time.Millisecond)
			// END DEBUG ================
			availableDaysIntersection := cu.IntSliceIntersection(daysAvailableCopy[userId], openDaysCopy)
			// DEBUG ================
			// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
			// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "9-%v", len(availableDaysIntersection)))
			// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
			// time.Sleep(25 * time.Millisecond)
			// END DEBUG ================
			randDay := availableDaysIntersection[utils.RandInt(0, len(availableDaysIntersection))]
			// DEBUG ================
			// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
			// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "Rand day - 10.1-%v", randDay))
			// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "10.2-%v", len(openSlotsCopy[randDay])))
			// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
			// time.Sleep(25 * time.Millisecond)
			// END DEBUG ================
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

		// DEBUG ================
		// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
		// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "13------\n%s", individual))
		// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
		// time.Sleep(100 * time.Millisecond)
		// END DEBUG ================

		// DEBUG ================
		// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
		// b, err = json.MarshalIndent(individual, "", "  ")
		// if err != nil {
		// 	logger.Error.Println(testutils.Scolourf(testutils.YELLOW, "8-%v", err))
		// }
		// logger.Error.Println(testutils.Scolourf(testutils.PURPLE, "8------\n%s", string(b)))
		// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
		// time.Sleep(100 * time.Millisecond)
		// END DEBUG ================
	}

	// DEBUG ================
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// logger.Error.Println(testutils.Scolour(testutils.GREEN, "Done"))
	// logger.Error.Println(testutils.Scolour(testutils.PURPLE, "============================="))
	// time.Sleep(100 * time.Millisecond)
	// END DEBUG ================

	return population
}
