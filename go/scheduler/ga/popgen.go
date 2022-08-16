package ga

import (
	cu "lib/collectionutils"
	"math/rand"
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

func populationGenerator(domain *Domain, popSize int) Individuals {
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

	// Add users as many times as they need to come into office
	usersToAdd := []string{}
	daysAvailable := make(map[string][]int, 0) // Used to indicate what days they are not already coming in
	for _, user := range domain.SchedulerData.Users {
		daysAvailable[*user.Id] = cu.SequentialSequence(0, 7)
		for i := 0; i < *user.OfficeDays; i++ { // TODO: @JonathanEnslin find out what do if no office days?
			usersToAdd = append(usersToAdd, *user.Id)
		}
	}

	// Remove days that users already come in from their available days
	for _, booking := range *schedulerData.CurrentBookings {
		bUserId := *booking.UserId
		if _, ok := daysAvailable[bUserId]; ok {
			// remove
			daysAvailable[bUserId] = cu.RemoveElementNoOrder(daysAvailable[bUserId], int(booking.GetWeekday()))
		}
	}

	// Used to maintain a copy of openSlots
	openSlotsCopy := make([][]int, len(openSlots))
	population := make([]*Individual, domain.Config.Size)
	for i := 0; i < domain.Config.Size; i++ {
		copy(openSlotsCopy, openSlots) // only a shallow copy is necessary
		individual := Individual{Gene: make([][]string, numDaysInWeek)}
		// Initialise individual to empty indiv
		for j := range individual.Gene {
			individual.Gene[j] = make([]string, slotSizes[j])
		}

		// Randomly assign employees to slots
		for _, userId := range usersToAdd {
			ri := rand.Intn(len(openSlotsCopy)) // rand index used to access open slots
			day, index := openSlotsCopy[ri][0], openSlotsCopy[ri][1]
			individual.Gene[day][index] = userId
			// Remove entry from openslots
			openSlots = append(openSlots[:ri], openSlots[ri+1:]...)
		}

		population[i] = &individual
	}
	return population
}

// func WeeklyPopulationGenerator(domain Domain, schedulerData data.SchedulerData, popSize int) Individuals {
// 	// Using representation Empl1 --- Empl2 --- Empl3 --- Empl1
// 	//						Mon       Mon       Thu       Tue
// 	//						Tue
// 	var population Individuals
// 	for j := 0; j < popSize; j++ {
// 		genes := []string{}
// 		for i := 0; i < len(schedulerData.Users); i++ {
// 			genes = append(genes, domain.Terminals[rand.Intn(len(domain.Terminals))])
// 		}
// 		population = append(population, &Individual{Gene: genes})
// 	}
// 	return population
// }
