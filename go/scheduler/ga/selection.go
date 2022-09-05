package ga

///////////////////////////////////////////////////
// WEEKLY

func WeeklyStubSelection(domain *Domain, individuals Individuals, count int) Individuals {
	return individuals[:count]
}

func WeeklyTournamentSelection(domain *Domain, individuals Individuals, count int) Individuals {
	var results Individuals
	for i := 0; i <= count; i++ {
		results = append(results, weeklyTournamentSelection(domain, individuals))
	}
	return results
}

func weeklyTournamentSelection(domain *Domain, individuals Individuals) *Individual {
	var tournament Individuals

	for i := 0; i <= domain.Config.TournamentSize; i++ {
		tournament = append(tournament, individuals.GetRandomIndividual().Clone())
	}

	var winner *Individual = tournament[0]

	for _, competitor := range tournament {
		if competitor.Fitness >= winner.Fitness {
			winner = competitor
		}
	}

	return winner
}

// SELECTION FUNCTION: FITNESS PROPORTIONATE SELECTION

///////////////////////////////////////////////////
// DAILY
