package ga

import (
	cu "lib/collectionutils"
	"lib/utils"
	"fmt"
	"strconv"
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
///////////////////////////////////////////////////

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
// General crossover code
///////////////////////////////////////////////////

//////////////////////
// Partially mapped crossover
//////////////////////

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

//////////////////////
// One-Point Crossover
//////////////////////
// Wikipedia

// OnePointCrossover is an invalid crossover that selects a random crossover point and crossovers everything to the right
func OnePointCrossover(domain *Domain, individuals Individuals, numOffspring int) Individuals {
	// Flatten the parents
	flatParent1, flatParent2 := cu.Flatten2DArr(individuals[0].Gene), cu.Flatten2DArr(individuals[1].Gene)

	// Get the crossover points
	xPoint := utils.RandInt(0, len(flatParent1))

	// Crossover
	offspring1, offspring2 := onePointCrossover(flatParent1, flatParent2, xPoint)

	// Calculate the original dimensions of the individuals
	sizes := make([]int, len(individuals[0].Gene))
	for i, col := range individuals[0].Gene {
		sizes[i] = len(col)
	}

	// child1 and child2 are the de-flattened offspring
	child1, child2 := cu.PartitionArray(offspring1, sizes), cu.PartitionArray(offspring2, sizes)
	return []*Individual{{child1, 0.0}, {child2, 0.0}}
}

func onePointCrossover[T comparable](p1, p2 []T, xP int) ([]T, []T) {
	if xP < 0  {
		xP = 0
	}
	// Make offspring arrays
	offspring1, offspring2 := make([]T, len(p1)), make([]T, len(p2))

	// Copy over start
	for i := 0; i < len(p1) && i < len(p2) && i < xP; i++ {
		offspring1[i] = p1[i]
		offspring2[i] = p2[i]
	}

	// CrossOver
	for i := xP; i < len(p1) && i < len(p2); i++ {
		offspring1[i] = p2[i]
		offspring2[i] = p1[i]
	}

	return offspring1, offspring2
}

//////////////////////
// Two-Point Crossover
//////////////////////
// Wikipedia

// TwoPointCrossover is an invalid crossover that crossovers everything between two crossover points
func TwoPointCrossover(domain *Domain, individuals Individuals, numOffspring int) Individuals {
	// Flatten the parents
	flatParent1, flatParent2 := cu.Flatten2DArr(individuals[0].Gene), cu.Flatten2DArr(individuals[1].Gene)

	// Get the crossover points
	xPoint1, xPoint2 := utils.RandInt(0, len(flatParent1)), utils.RandInt(0, len(flatParent1))

	// Ensure that the first crossover point is not larger than the second
	if xPoint1 > xPoint2 {
		xPoint1, xPoint2 = xPoint2, xPoint1
	}

	// Crossover
	offspring1, offspring2 := twoPointCrossover(flatParent1, flatParent2, xPoint1, xPoint2)

	// Calculate the original dimensions of the individuals
	sizes := make([]int, len(individuals[0].Gene))
	for i, col := range individuals[0].Gene {
		sizes[i] = len(col)
	}

	// child1 and child2 are the de-flattened offspring
	child1, child2 := cu.PartitionArray(offspring1, sizes), cu.PartitionArray(offspring2, sizes)
	return []*Individual{{child1, 0.0}, {child2, 0.0}}
}

func twoPointCrossover[T comparable](p1, p2 []T, xP1, xP2 int) ([]T, []T) {
	if xP1 < 0  {
		xP1 = 0
	}
	if xP2 < 0  {
		xP2 = 0
	}

	// Make offspring arrays
	offspring1, offspring2 := make([]T, len(p1)), make([]T, len(p2))

	// Copy over start
	for i := 0; i < len(p1) && i < len(p2) && i < xP1; i++ {
		offspring1[i] = p1[i]
		offspring2[i] = p2[i]
	}

	// CrossOver
	for i := xP1; i < len(p1) && i < len(p2) && i < xP2; i++ {
		offspring1[i] = p2[i]
		offspring2[i] = p1[i]
	}

	// Copy over end
	for i := xP2; i < len(p1) && i < len(p2); i++ {
		offspring1[i] = p1[i]
		offspring2[i] = p2[i]
	}

	return offspring1, offspring2
}

//////////////////////
// Cycle Crossover
//////////////////////
// https://www.researchgate.net/figure/Cycle-crossover-CX_fig2_226665831#:~:text=Cycle%20crossover%20(CX)%20The%20cycle,from%20one%20of%20the%20parents.
// The cycle crossover operator(Figure 3) was proposed by Oliver et al. (1987)
// Can only be used in daily scheduler

// CycleCrossover is a valid crossover method
func CycleCrossover(domain *Domain, individuals Individuals, numOffspring int) Individuals {
	// Flatten the parents
	flatParent1, flatParent2 := cu.Flatten2DArr(individuals[0].Gene), cu.Flatten2DArr(individuals[1].Gene)

	// Resource map
	resources := make(map[string]int)
	for i, resource := range domain.SchedulerData.Resources {
		resources[*resource.Id] = i
	}

	// Resource reverse map
	reverseResources := make(map[int]string)
	for key, value := range reverseResources {
		resources[value] = key
	}

	// convert parents to the right representation
	var intParent1 []int
	var intParent2 []int
	for _, resource := range flatParent1 {
		intParent1 = append(intParent1, resources[resource])
	}
	for _, resource := range flatParent2 {
		intParent2 = append(intParent2, resources[resource])
	}
	for i := 0; i < len(domain.SchedulerData.Resources); i++ {
		if !contains(flatParent1, fmt.Sprint(i)) {
			intParent1 = append(intParent1, i)
		}
	}
	for i := 0; i < len(domain.SchedulerData.Resources); i++ {
		if !contains(flatParent2, fmt.Sprint(i)) {
			intParent2 = append(intParent2, i)
		}
	}

	// Crossover
	offspring1, offspring2 := cycleCrossover(intParent1, intParent2)
	offspring1 = offspring1[:len(flatParent1) - 1]
	offspring2 = offspring2[:len(flatParent2) - 1]

	// Convert int array to string array
	var result1 []string
	for _, i := range offspring1 {
		result1 = append(result1, strconv.Itoa(i))
	}

	var result2 []string
	for _, i := range offspring2 {
		result2 = append(result2, strconv.Itoa(i))
	}

	// Calculate the original dimensions of the individuals
	sizes := make([]int, len(individuals[0].Gene))
	for i, col := range individuals[0].Gene {
		sizes[i] = len(col)
	}

	// child1 and child2 are the de-flattened offspring
	child1, child2 := cu.PartitionArray(result1, sizes), cu.PartitionArray(result2, sizes)
	return []*Individual{{child1, 0.0}, {child2, 0.0}}
}

func cycleCrossover(p1, p2 []int) ([]int, []int) {
	// Detect and swap odd number cycles
	visited := make(map[int]bool)
	oddCycle := false
	for i := 0; i < len(p1); i++ {
		// if not already part of a cycle
		if !visited[i] {
			var cycle []int
			cycle = append(cycle, i)
			visited[i] = true
			
			current := p2[i]
			odd := false
			for current != i {
				if odd {
					temp:= p2[current]
					current = temp
				} else {
					temp:= p1[current]
					cycle = append(cycle, temp)
					current = temp
					visited[temp] = true
				}
				odd = !odd
			}

			if oddCycle {
				for _, index := range cycle {
					temp := p1[index]
					p1[index] = p2[index]
					p2[index] = temp
				}
			}

			oddCycle = !oddCycle
		}
	}

	return p1, p2
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}