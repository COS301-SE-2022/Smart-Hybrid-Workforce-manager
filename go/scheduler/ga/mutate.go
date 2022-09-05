package ga

import (
	"lib/utils"
)

///////////////////////////////////////////////////
// WEEKLY

func WeeklyStubMutate(domain *Domain, individuals Individuals) Individuals {
	return individuals.ClonePopulation()
}

// Performs of a version of swap mutation, where users are shifted around between days, but are kept valid
func WeeklyWeeklyDayVResourceMutateSwapValid(domain *Domain, individuals Individuals) Individuals {
	var results Individuals
	for _, individual := range individuals {
		copiedIndiv := individual.Clone()
		validSwap(copiedIndiv, len(domain.Terminals)/7) // semi-random
		results = append(results, copiedIndiv)
	}
	return results
}

// Swap random users to differnt days, but keep the individual valid
func validSwap(indiv *Individual, mutationDegree int) *Individual {
	// An array of maps, where each map contains the counts that users come in on a certain day
	// the index of the day corresponds to the day
	// map[user id]# of times a users id appears on that day
	usersComingInOnDay := make([]map[string]bool, len(indiv.Gene))
	// initialise map
	for i := 0; i < len(usersComingInOnDay); i++ {
		usersComingInOnDay[i] = make(map[string]bool)
		for _, userid := range indiv.Gene[i] {
			if userid != "" {
				usersComingInOnDay[i][userid] = false
			}
		}
	}
	// Perform mutations
	for i := 0; i < mutationDegree; i++ {
		dayi1, sloti1 := randSlot(indiv)
		// select a day that is not the same day as dayi1
		dayi2, sloti2 := randSlot(indiv)
		for dayi2 == dayi1 && dayi2 != -1 { // if dayi2 == -1 it means that there might no be enough slots
			dayi2, sloti2 = randSlot(indiv)
		}

		if dayi1 == -1 || dayi2 == -1 { // Mutation can not be performed
			break
		}

		_, contains1 := usersComingInOnDay[dayi1][indiv.Gene[dayi2][sloti2]]
		_, contains2 := usersComingInOnDay[dayi2][indiv.Gene[dayi1][sloti1]]
		if !contains1 && !contains2 {
			performSwap(indiv, dayi1, dayi2, sloti1, sloti2)
		}
	}
	return indiv
}

func performSwap(indiv *Individual, day1, day2, sloti1, sloti2 int) {
	indiv.Gene[day1][sloti1], indiv.Gene[day2][sloti2] =
		indiv.Gene[day2][sloti2], indiv.Gene[day1][sloti1]
}

// Returns the day and index of a random slot in an individual
func randSlot(indiv *Individual) (randDay int, randSlot int) {
	randDay = utils.RandInt(0, len(indiv.Gene)) // First mutation point
	if len(indiv.Gene[randDay]) > 0 {           // Will panic if <= 0
		randSlot = utils.RandInt(0, len(indiv.Gene[randDay]))
	} else {
		randSlot = -1 // Use -1 to indicate that no slots are available
	}
	return
}

func WeeklyDayVResourceMutateSwap(domain *Domain, individuals Individuals) Individuals {
	var results Individuals
	// Mutate each individual
	for _, individual := range individuals {
		copiedIndividual := individual.Clone()
		// Pick random slot, swap with another random slot
		// randDay1 := rand.Intn(len(copiedIndividual.Gene))        // First mutation point
		randDay1 := utils.RandInt(0, len(copiedIndividual.Gene)) // First mutation point
		randDay2 := utils.RandInt(0, len(copiedIndividual.Gene))

		randSloti1 := utils.RandInt(0, len(copiedIndividual.Gene[randDay1]))
		randSloti2 := utils.RandInt(0, len(copiedIndividual.Gene[randDay2]))

		copiedIndividual.Gene[randDay1][randSloti1], copiedIndividual.Gene[randDay2][randSloti2] =
			copiedIndividual.Gene[randDay2][randSloti2], copiedIndividual.Gene[randDay1][randSloti1]
		results = append(results, copiedIndividual)
	}
	return results
}

func WeeklyDayVResouceMutate(domain *Domain, individuals Individuals) Individuals {
	var results Individuals
	for _, individual := range individuals {
		copiedIndividual := individual.Clone()
		if len(domain.Terminals) == 0 {
			continue // No mutation of this type can be done if there are no terminals
		}
		for i := range individual.Gene {
			// randDay1 := rand.Intn(len(individual.Gene)) // First mutation point
			if len(copiedIndividual.Gene[i]) == 0 { // No mutation can be performed if there are no slots at all
				continue
			}
			randSloti := utils.RandInt(0, len(copiedIndividual.Gene[i]))
			// mutate everything
			for j := randSloti; j < len(copiedIndividual.Gene[i]); j++ {
				// TODO set chance to be empty
				if utils.RandInt(0, 100) > 4 {
					copiedIndividual.Gene[i][j] = ""
				} else {
					copiedIndividual.Gene[i][j] = domain.GetRandomTerminal()
				}
			}
		}
		results = append(results, copiedIndividual)
	}
	return results
}

///////////////////////////////////////////////////
// DAILY
