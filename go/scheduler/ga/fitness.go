package ga

func StubFitness(domain *Domain, individuals Individuals) []float64 {
	var result []float64
	for i := 0; i < len(individuals); i++ {
		result = append(result, 0.0)
		individuals[i].Fitness = 0.0
	}
	return result
}
