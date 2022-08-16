package ga

import "math/rand"

func StubCrossOver(domain *Domain, individuals Individuals, selectionFunc Selection, fitnii []float64, offspring int) Individuals {
	if len(individuals) == 0 || offspring == 0 {
		return nil
	}
	var result Individuals
	for i := 0; i < offspring; i++ {
		result = append(result, individuals[i%len(individuals)].Clone())
	}
	return result
}

func GenericSinglePointCrossover(domain *Domain, individuals Individuals, selectionFunc Selection, fitnii []float64, offspring int) Individuals {
	if len(individuals) == 0 || offspring == 0 {
		return nil
	}
	var result Individuals
	for i := 0; i < offspring; i++ {
		rand.Intn(len(individuals[i].Gene))
		result = append(result, individuals[i%len(individuals)].Clone())
	}
	return result
}
