package ga

import (
	. "lib/collectionutils"
	"lib/utils"
)

///////////////////////////////////////////////////
// WEEKLY

func WeeklyStubCrossOver(domain *Domain, individuals Individuals, selectionFunc Selection, offspring int) Individuals {
	if len(individuals) == 0 || offspring == 0 {
		return nil
	}
	var result Individuals
	for i := 0; i < offspring; i++ {
		result = append(result, individuals[i%len(individuals)].Clone())
	}
	return result
}

func WeeklyDayVResourceCrossover(domain *Domain, individuals Individuals, selectionFunc Selection, offspring int) Individuals {
	var results Individuals
	for i := 0; i <= offspring; i++ {
		// select parents
		parents := selectionFunc(domain, individuals, 2)
		results = append(results, weeklyDayVResourceCrossover(domain, parents, 2)...)
	}
	return results
}

func weeklyDayVResourceCrossover(domain *Domain, individuals Individuals, offspring int) Individuals {
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

// A valid crossover, that works similarly to PMX 2-point crossover
// It initially flattens an individual, and then performs crossover on the flattened crossover
func weeklyFlattenCrossoverValid(domain *Domain, individuals Individuals, offspring int) Individuals {
	flatParent1, flatParent2 := Flatten2DArr(individuals[0].Gene), Flatten2DArr(individuals[0].Gene)

	xPoint1, xPoint2 := utils.RandInt(0, len(flatParent1)), utils.RandInt(0, len(flatParent1))

	if xPoint1 > xPoint2 {
		xPoint1, xPoint2 = xPoint2, xPoint1
	}

	flatChild1, flatChild2 := twoPointSwap(flatParent1, flatParent2, xPoint1, xPoint2)

	// TODO: @JonathanEnslin - Make individuals valid

	sizes := make([]int, len(individuals[0].Gene))
	for i, col := range individuals[0].Gene {
		sizes[i] = len(col)
	}

	child1, child2 := PartitionArray(flatChild1, sizes), PartitionArray(flatChild2, sizes)
	return []*Individual{{child1, 0.0}, {child2, 0.0}}
}

func twoPointSwap(arr1, arr2 []string, xP1, xP2 int) ([]string, []string) {
	res1, res2 := make([]string, len(arr1)), make([]string, len(arr1))
	for i := 0; i < xP1; i++ {
		res1[i] = arr1[i]
		res2[i] = arr2[i]
	}
	for i := xP1; i < xP2; i++ {
		res1[i] = arr2[i]
		res2[i] = arr1[i]
	}
	for i := xP2; i < len(arr1); i++ {
		res1[i] = arr1[i]
		res2[i] = arr2[i]
	}
	return res1, res2
}

///////////////////////////////////////////////////
// DAILY
