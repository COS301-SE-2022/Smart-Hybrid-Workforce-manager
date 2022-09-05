package ga

import "lib/utils"

func StubCrossOver(domain *Domain, individuals Individuals, selectionFunc Selection, offspring int) Individuals {
	if len(individuals) == 0 || offspring == 0 {
		return nil
	}
	var result Individuals
	for i := 0; i < offspring; i++ {
		result = append(result, individuals[i%len(individuals)].Clone())
	}
	return result
}

// Not complete
// func GenericSinglePointCrossover(domain *Domain, individuals Individuals, selectionFunc Selection, offspring int) Individuals {
// 	if len(individuals) == 0 || offspring == 0 {
// 		return nil
// 	}
// 	var result Individuals
// 	for i := 0; i < offspring; i++ {
// 		rand.Intn(len(individuals[i].Gene))
// 		result = append(result, individuals[i%len(individuals)].Clone())
// 	}
// 	return result
// }

func DayVResourceCrossover(domain *Domain, individuals Individuals, selectionFunc Selection, offspring int) Individuals {
	var results Individuals
	for i := 0; i <= offspring; i++ {
		// select parents
		parents := selectionFunc(domain, individuals, 2)
		results = append(results, dayVResourceCrossover(domain, parents, 2)...)
	}
	return results
}

func dayVResourceCrossover(domain *Domain, individuals Individuals, offspring int) Individuals {
	if len(individuals) != 2 {
		return nil
	}
	parent1 := individuals[0].Clone()
	parent2 := individuals[1].Clone()

	// pick random gene to crossover
	crossoverParent1 := parent1.Gene[utils.RandInt(0, len(parent1.Gene))]
	crossoverParent2 := parent2.Gene[utils.RandInt(0, len(parent2.Gene))]

	// pick random slot to crossover
	crossoverParent1Slot := utils.RandInt(0, len(crossoverParent1))
	crossoverParent2Slot := utils.RandInt(0, len(crossoverParent2))

	// crossover with fixed length
	length := utils.RandInt(0, len(crossoverParent1))

	// crossover for length
	for i := 0; i <= length; i++ {
		// no overflow
		if len(crossoverParent1) < crossoverParent1Slot+i && len(crossoverParent2) < crossoverParent2Slot+i {
			temp := crossoverParent1[crossoverParent1Slot+i]
			crossoverParent1[crossoverParent1Slot+i] = crossoverParent2[crossoverParent2Slot+i]
			crossoverParent2[crossoverParent2Slot+i] = temp
		}
	}

	return Individuals{parent1, parent2}
}
