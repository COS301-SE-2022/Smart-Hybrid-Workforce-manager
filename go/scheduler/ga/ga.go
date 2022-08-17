package ga

import (
	"fmt"
	"lib/utils"
	"math/rand"
	"scheduler/data"
	"strings"
)

// Individuals is a collection of solutions to the domain problem
type Individuals []*Individual

type Crossover func(domain *Domain, individuals Individuals, selectionFunc Selection, offspring int) Individuals
type Fitness func(domain *Domain, individuals Individuals) []float64
type Mutate func(domain *Domain, individuals Individuals) Individuals
type Selection func(domain *Domain, individuals Individuals, count int) Individuals
type PopulationGenerator func(domain *Domain, popSize int) Individuals

// Domain represents domain information to the problem
type Domain struct {
	// Weekly scheduler: User array with user id's duplicated for amount per week
	Terminals []string

	Config        *data.Config
	SchedulerData *data.SchedulerData
}

func (domain *Domain) GetRandomTerminal() string {
	return domain.Terminals[utils.RandInt(0, len(domain.Terminals))]
}

// Individual represents a solution to the domain problem
type Individual struct {
	// Weekly scheduler: Day of week x Resource count
	Gene    [][]string
	Fitness float64
}

// Clone clones an individual
func (individual *Individual) Clone() *Individual {
	newGene := make([][]string, len(individual.Gene))
	for i := 0; i < len(individual.Gene); i++ {
		newGene[i] = make([]string, len(individual.Gene[i]))
		copy(newGene[i], individual.Gene[i])
	}
	newIndividual := &Individual{Gene: newGene, Fitness: individual.Fitness}
	return newIndividual
}

// ClonePopulation clones an entire population
func (population Individuals) ClonePopulation() []*Individual {
	cloned := make([]*Individual, len(population))
	for _, individual := range population {
		cloned = append(cloned, individual.Clone())
	}
	return cloned
}

// GetRandomIndividual retrieves a random individual in the population
func (population Individuals) GetRandomIndividual() *Individual {
	return population[utils.RandInt(0, len(population))]
}

// GA is a generic configurable genetic algorithm that produces multiple solutions to the domain problem
func GA(domain Domain, crossover Crossover, fitness Fitness, mutate Mutate, selection Selection, populationGenerator PopulationGenerator) Individuals {
	// Seed
	rand.Seed(int64(domain.Config.Seed))

	// Create initial pop and calculate fitnesses
	population := populationGenerator(&domain, 5)
	fitness(&domain, population) // TODO Change to useful individuals

	// Run ga
	stoppingCondition := true
	for i := 0; i < domain.Config.Generations && stoppingCondition; i++ {
		crossOverAmount := (domain.Config.PopulationSize * int(domain.Config.CrossOverRate))
		mutateAmount := (domain.Config.PopulationSize * int(domain.Config.MutationRate))
		carryAmount := domain.Config.PopulationSize - crossOverAmount - mutateAmount // TODO: Find out Anna if is guicci

		// evolve
		individualsOffspring := crossover(&domain, population, selection, crossOverAmount)
		individualsMutated := mutate(&domain, selection(&domain, population, mutateAmount))
		individualsCarry := selection(&domain, population, carryAmount)

		population := append(individualsOffspring, individualsMutated...)
		population = append(population, individualsCarry...)

		fitness(&domain, population)
	}

	return population
}

//       Monday   -   Tuesday   -  Wednesday
// 08:00 emp1, emp2
// 09:00
// 10:10

// emp1     emp2    emp3
// Mon		Mon
// 0-1		0-1
func (individual Individual) String() string {
	// Returns table representation of an individual
	// 	userIds := make(map[string]int)
	// 	for dayi := range individual.Gene {
	// 		for _, userId := range individual.Gene[dayi] {
	// 			if _, ok := userIds[userId]; !ok && !(userId == "") {
	// 				userIds[userId] = len(userIds)
	// 			}
	// 		}
	// 	}
	maxSlotsize := -1
	for _, day := range individual.Gene {
		if len(day) > maxSlotsize {
			maxSlotsize = len(day)
		}
	}
	table := make([][]string, len(individual.Gene))
	for i := range table {
		table[i] = make([]string, maxSlotsize+3) // +1 for day thing + 1 for ending border thingy
	}

	for i := 0; i < len(individual.Gene); i++ {
		for j := 0; j < len(table[i]); j++ {
			table[i][j] = strings.Repeat(" ", 39)
		}
	}

	for i := range individual.Gene {
		table[i][0] = fmt.Sprintf("%38s|", fmt.Sprintf("Day-%d", i))
		table[i][1] = strings.Repeat("=", 39)
	}

	for i := 0; i < len(individual.Gene); i++ {
		counter := 0
		for j := 0; j < len(individual.Gene[i]); j++ {
			counter = j
			table[i][j+2] = individual.Gene[i][j]
			table[i][j+2] = fmt.Sprintf("%38s|", individual.Gene[i][j])

		}
		table[i][counter+3] = strings.Repeat("-", 38) + "|"
	}

	// More borders
	for i := 1; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i-1][j] == strings.Repeat(" ", 39) && table[i][j] != strings.Repeat(" ", 39) {
				table[i-1][j] = fmt.Sprintf("%38s|", "")
			}
		}
	}

	tableStr := ""
	// for key, val := range userIds {
	// 	tableStr += fmt.Sprintf("%s => %d\n", key, val)
	// }
	// tableStr += "\n"

	for j := 0; j < len(table[0]); j++ {
		for i := 0; i < len(table); i++ {
			tableStr += table[i][j]
		}
		tableStr += "\n"
	}
	return tableStr
}
