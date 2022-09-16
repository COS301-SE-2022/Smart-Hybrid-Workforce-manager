package ga

import (
	cu "lib/collectionutils"
	"lib/utils"
)

// Function of this type gets passed into CrossoverCaller
type CrossoverOperator func(domain *Domain, individuals Individuals, numOffspring int) Individuals

// CrossoverCaller gets passed to the GA, it uses the specified operator to perform crossover
func CrossoverCaller(crossoverOperator CrossoverOperator, domain *Domain, individuals Individuals, selectionFunc Selection, offspring int) Individuals {
	var results Individuals
	for i := 0; i <= offspring; i++ {
		// select parents
		parents := selectionFunc(domain, individuals, 2)
		results = append(results, crossoverOperator(domain, parents, 2)...)
	}
	return results[:offspring]
}

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
		if len(crossoverParent1) > crossoverParent1Slot+i && len(crossoverParent2) > crossoverParent2Slot+i {
			temp := crossoverParent1[crossoverParent1Slot+i]
			crossoverParent1[crossoverParent1Slot+i] = crossoverParent2[crossoverParent2Slot+i]
			crossoverParent2[crossoverParent2Slot+i] = temp
		}
	}

	return Individuals{parent1, parent2}
}

///////////////////////////////////////////////////
// DAILY

///////////////////////////////////////////////////
// General crossover code

// A valid (daily) crossover, that works similarly to PMX 2-point crossover
// It initially flattens an individual, and then performs pmx crossover on the flattened crossover
func PartiallyMappedFlattenCrossoverValid(domain *Domain, individuals Individuals, numOffspring int) Individuals {
	// Flatten the parents
	flatParent1, flatParent2 := cu.Flatten2DArr(individuals[0].Gene), cu.Flatten2DArr(individuals[1].Gene)

	// Get the crossover points
	xPoint1, xPoint2 := utils.RandInt(0, len(flatParent1)), utils.RandInt(0, len(flatParent1))

	// Ensure that the first crossover point is not larger than the second
	if xPoint1 > xPoint2 {
		xPoint1, xPoint2 = xPoint2, xPoint1
	}

	offspring1, offspring2 := PMX(flatParent1, flatParent2, xPoint1, xPoint2)

	// Calculate the original dimensions of the individuals
	sizes := make([]int, len(individuals[0].Gene))
	for i, col := range individuals[0].Gene {
		sizes[i] = len(col)
	}

	// child1 and child2 are the de-flattened offspring
	child1, child2 := cu.PartitionArray(offspring1, sizes), cu.PartitionArray(offspring2, sizes)
	return []*Individual{{child1, 0.0}, {child2, 0.0}}
}

// twoPointSwap swaps the elements from xP1 up and to excluding xP2 between arr1 and arr2
func twoPointSwap[T any](arr1, arr2 []T, xP1, xP2 int) ([]T, []T) {
	res1, res2 := make([]T, len(arr1)), make([]T, len(arr1))
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

// PMX performs partially mapped crossover, where p1 and p2 are parents, and xP1 and xP2 are
// the crossover points
func PMX[T comparable](p1, p2 []T, xP1, xP2 int) ([]T, []T) {
	// Make offspring arrays
	offspring1, offspring2 := make([]T, len(p1)), make([]T, len(p2))

	// Map the genetic material inside the crossover section
	p1XSection, p2XSection := make(map[T]int), make(map[T]int)
	for i := xP1; i < xP2; i++ {
		p1XSection[p1[i]] = i
		p2XSection[p2[i]] = i
	}

	// Cross over the crossover sections
	for i := xP1; i < xP2; i++ {
		offspring1[i] = p2[i]
		offspring2[i] = p1[i]
	}

	// Fill in the rest of the chromosomes by mapping where necessary
	for i := 0; i < xP1; i++ {
		offspring1[i] = FindValid(i, p1, p2XSection)
		offspring2[i] = FindValid(i, p2, p1XSection)
	}

	for i := xP2; i < len(p1); i++ {
		offspring1[i] = FindValid(i, p1, p2XSection)
		offspring2[i] = FindValid(i, p2, p1XSection)
	}
	return offspring1, offspring2
}

func FindValid[T comparable](index int, parent []T, otherParentMap map[T]int) T {
	// Check if element about to be copied from parent is already present inside the crossed over section
	if !cu.MapHasKey(otherParentMap, parent[index]) {
		return parent[index] // If it isn't, return that element since it can be inserted
	} else {
		return FindValid(otherParentMap[parent[index]], parent, otherParentMap) // If it is, find and return a valid element
	}
}
