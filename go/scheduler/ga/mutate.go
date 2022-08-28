package ga

import (
	"lib/utils"
)

func StubMutate(domain *Domain, individuals Individuals) Individuals {
	return individuals.ClonePopulation()
}

func DayVResourceMutateSwap(domain *Domain, individuals Individuals) Individuals {
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

func DayVResouceMutate(domain *Domain, individuals Individuals) Individuals {
	var results Individuals
	for _, individual := range individuals {
		copiedIndividual := individual.Clone()
		for i := range individual.Gene {
			// randDay1 := rand.Intn(len(individual.Gene)) // First mutation point
			randSloti := utils.RandInt(0, len(copiedIndividual.Gene[i]))
			// mutate everything
			for j := randSloti; j < len(copiedIndividual.Gene[i]); j++ {
				// TODO set chance to be empty
				if utils.RandInt(0, 100) > 50 {
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
