package ga

type Comparator func(a, b float64) bool

///////////////////////////////////////////////////
// WEEKLY

// Returns a >= b
func GTorEq(a, b float64) bool {
	return a >= b
}

// Returns a < b
func LT(a, b float64) bool {
	return a < b
}

func WeeklyStubSelection(domain *Domain, individuals Individuals, count int) Individuals {
	return individuals[:count]
}

func WeeklyTournamentSelectionFitness(domain *Domain, individuals Individuals, count int) Individuals {
	return WeeklyTournamentSelection(domain, individuals, count, GTorEq)
}

func WeeklyTournamentSelectionCost(domain *Domain, individuals Individuals, count int) Individuals {
	return WeeklyTournamentSelection(domain, individuals, count, LT)
}

func WeeklyTournamentSelection(domain *Domain, individuals Individuals, count int, comparator Comparator) Individuals {
	var results Individuals
	for i := 0; i < count; i++ {
		results = append(results, weeklyTournamentSelection(domain, individuals, comparator))
	}
	return results
}

func weeklyTournamentSelection(domain *Domain, individuals Individuals, comparator Comparator) *Individual {
	var tournament Individuals

	for i := 0; i <= domain.Config.TournamentSize; i++ {
		tournament = append(tournament, individuals.GetRandomIndividual().Clone())
	}

	var winner *Individual = tournament[0]

	for _, competitor := range tournament {
		// if competitor.Fitness >= winner.Fitness {
		if comparator(competitor.Fitness, winner.Fitness) {
			// if winner.Fitness <= competitor.Fitness {
			winner = competitor
		}
	}

	return winner
}

// SELECTION FUNCTION: FITNESS PROPORTIONATE SELECTION

///////////////////////////////////////////////////
// DAILY
