package ga

import (
	cu "lib/collectionutils"
	"lib/utils"
)

///////////////////////////////////////////////////
// WEEKLY

func WeeklyStubMutate(domain *Domain, individuals Individuals) Individuals {
	return individuals.ClonePopulation()
}

// Performs of a version of swap mutation, where users are shifted around between days, but are kept valid
func WeeklyDayVResourceMutateSwapValid(domain *Domain, individuals Individuals) Individuals {
	var results Individuals
	for _, individual := range individuals {
		copiedIndiv := individual.Clone()
		validWeeklySwap(copiedIndiv, len(domain.Terminals)/8) // semi-random
		results = append(results, copiedIndiv)
	}
	return results
}

// Swap random users to differnt days, but keep the individual valid
func validWeeklySwap(indiv *Individual, mutationDegree int) *Individual {
	// An array of maps, where each map contains the index that users come in on a certain day
	// the index of the day corresponds to the day
	// map[user id]index that a user comes in
	usersComingInOnDay := make([]map[string]int, len(indiv.Gene))
	// Keeps track of which indices on specific days have assigned user ids
	indicesWithUsers := make([][]int, len(indiv.Gene))
	for i := 0; i < len(indicesWithUsers); i++ {
		indicesWithUsers[i] = []int{}
	}
	// initialise map and indicesWithUsersArray
	for i := 0; i < len(usersComingInOnDay); i++ {
		usersComingInOnDay[i] = make(map[string]int)
		for j, userid := range indiv.Gene[i] {
			if userid != "" {
				indicesWithUsers[i] = append(indicesWithUsers[i], j)
				usersComingInOnDay[i][userid] = j
			}
		}
	}

	// Perform mutations
	for i := 0; i < mutationDegree; i++ {
		// dayi1, sloti1 := randSlot(indiv)
		// Get a random user to swap around
		// Search to see which days have assigned users
		daysWithUsers := indicesOfNonEmptyArrs(indicesWithUsers) // Get days that actually have users
		dayi1, sloti1 := -1, -1
		if len(daysWithUsers) > 0 {
			dayi1 = daysWithUsers[utils.RandInt(0, len(daysWithUsers))]                      // Select a random day
			sloti1 = indicesWithUsers[dayi1][utils.RandInt(0, len(indicesWithUsers[dayi1]))] // Get a slot with a user in it
		}

		if dayi1 == -1 || sloti1 == -1 {
			continue
		}

		// select a day that is not the same day as dayi1
		dayi2, sloti2 := randSlot(indiv)
		for dayi2 == dayi1 && dayi2 != -1 && sloti2 != -1 { // if dayi2 == -1 it means that there might no be enough slots
			dayi2, sloti2 = randSlot(indiv)
		}

		if sloti2 == -1 || dayi2 == -1 { // Mutation can not be performed
			continue
		}

		// Swap if allowed
		if !cu.MapHasKey(usersComingInOnDay[dayi2], indiv.Gene[dayi1][sloti1]) && !cu.MapHasKey(usersComingInOnDay[dayi1], indiv.Gene[dayi2][sloti2]) { // Only swap if user is not already in that day
			oldDayi1Id := indiv.Gene[dayi1][sloti1]
			oldDayi2Id := indiv.Gene[dayi2][sloti2]
			performSwap(indiv, dayi1, dayi2, sloti1, sloti2)
			// Update maps and arrays
			usersComingInOnDay[dayi2][oldDayi1Id] = sloti2
			delete(usersComingInOnDay[dayi1], oldDayi1Id)

			// If user was swapped with another, update for dayi1
			if indiv.Gene[dayi1][sloti1] != "" {
				usersComingInOnDay[dayi1][oldDayi2Id] = sloti1
				delete(usersComingInOnDay[dayi2], oldDayi2Id)
			} else {
				cu.RemElemenAtI(indicesWithUsers[dayi1], sloti1)
				indicesWithUsers[dayi2] = append(indicesWithUsers[dayi2], sloti2)
			}
		}

		// _, contains1 := usersComingInOnDay[dayi1][indiv.Gene[dayi2][sloti2]]
		// _, contains2 := usersComingInOnDay[dayi2][indiv.Gene[dayi1][sloti1]]
		// if !contains1 && !contains2 {
		// 	performSwap(indiv, dayi1, dayi2, sloti1, sloti2)
		// }
	}
	return indiv
}

func indicesOfNonEmptyArrs[T any](arr2D [][]T) []int {
	indices := []int{}
	for i, inner := range arr2D {
		if len(inner) > 0 {
			indices = append(indices, i)
		}
	}
	return indices
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
				if utils.RandInt(0, 100) > 10 {
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

func DailyMutate(domain *Domain, individuals Individuals) Individuals {
	var results Individuals
	for _, individual := range individuals {
		copiedIndividual := individual.Clone()
		if len(individual.Gene) == 0 {
			return nil
		}
		slot1 := utils.RandInt(0, len(copiedIndividual.Gene[0]))
		slot2 := utils.RandInt(0, len(copiedIndividual.Gene[0]))
		if slot1 > slot2 {
			temp := slot1
			slot1 = slot2
			slot2 = temp
		}
		for i := slot1; i < slot2; i++ {
			copiedIndividual.Gene[0][i] = domain.GetRandomTerminal()
		}
		results = append(results, copiedIndividual)
	}
	return results
}

// DailyMutateValid produces len(individuals) amount of new mutated individuals
func DailyMutateValid(domain *Domain, individuals Individuals) Individuals {
	var results Individuals
	for _, indiv := range individuals {
		results = append(results, dailyMutateValid(domain, indiv, 0.5, len(indiv.Gene[0])/12, len(indiv.Gene[0])/12))
	}
	return results
}

// Performs mutation on a daily individual
// with a certain probability it will either swap around resources already
// in an individual, or pull new resources from the terminals (if there are new resources)
func dailyMutateValid(domain *Domain, individual *Individual, swapProbability float64, swapAmount, pullAmount int) *Individual {
	decider := utils.RandFloat64()
	if decider <= swapProbability {
		// Apply swap mutate
		return dailySwapMutate(domain, individual, swapAmount)
	}
	// Attempt pull mutate
	mutated := dailyPullMutateValid(domain, individual, pullAmount)
	if mutated == nil { // Apply swap mutate if pull mutate could not be performed
		return dailySwapMutate(domain, individual, swapAmount)
	}
	return mutated
}

// Swaps around resources inside a daily individual, it does this swapAmount times
// Individuals will remain valid if they were valid to begin with\
func dailySwapMutate(domain *Domain, individual *Individual, swapAmount int) *Individual {
	copiedIndividual := individual.Clone()
	gene := copiedIndividual.Gene
	if len(gene[0]) == 0 {
		return copiedIndividual // No users to be assigned resources
	}
	for i := 0; i < swapAmount; i++ {
		// Find random indices to swap inside the individaul
		randi1, randi2 := utils.RandInt(0, len(gene[0])), utils.RandInt(0, len(gene[0]))
		// Swap the resources
		gene[0][randi1], gene[0][randi2] = gene[0][randi2], gene[0][randi1]
	}
	return copiedIndividual
}

// dailyPullMutateValid will pull new resources at random from the terminals, and
// replace assigned resources at random, while keeping the individual valid, returns nil
// if no terminals can be taken
func dailyPullMutateValid(domain *Domain, individual *Individual, pullAmount int) *Individual {
	copiedIndividual := individual.Clone()
	gene := copiedIndividual.Gene
	availableTerminals := cu.SliceDifference(domain.Terminals, gene[0])
	if len(gene[0]) == 0 {
		return copiedIndividual // no users to be assigned resoures
	}
	if len(availableTerminals) == 0 {
		return nil // No terminals
	}
	for i := 0; i < pullAmount; i++ {
		// Get index for a random resource from the available resources
		rTerminali := utils.RandInt(0, len(availableTerminals))
		// Get index for a random user/resource in the gene
		rGenei := utils.RandInt(0, len(gene[0]))
		// Exchange the resource assigned to the individual with the new resource obtained from the available terminals
		gene[0][rGenei], availableTerminals[rTerminali] = availableTerminals[rTerminali], gene[0][rGenei]
	}
	return copiedIndividual
}
