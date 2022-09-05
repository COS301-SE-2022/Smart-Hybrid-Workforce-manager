package ga

func StubSelection(domain *Domain, individuals Individuals, count int) Individuals {
	return individuals[:count]
}

func TournamentSelection(domain *Domain, individuals Individuals, count int) Individuals {
	var results Individuals
	for i := 0; i <= count; i++ {
		results = append(results, tournamentSelection(domain, individuals))
	}
	return results
}

func tournamentSelection(domain *Domain, individuals Individuals) *Individual {
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
