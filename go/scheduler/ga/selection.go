package ga

func StubSelection(domain Domain, individuals Individuals, fitness []float64, count int) Individuals {
	return individuals[:count]
}

func TournamentSelection(domain Domain, individuals Individuals, fitness []float64, count int) Individuals {
	return nil
}
