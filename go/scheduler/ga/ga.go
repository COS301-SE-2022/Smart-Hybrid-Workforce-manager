package ga

import (
	"lib/utils"
	"math/rand"
	"scheduler/data"
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
	newIndividual := &Individual{Gene: individual.Gene}
	return newIndividual
}

// ClonePopulation clones an entire population
func (population Individuals) ClonePopulation() []*Individual {
	cloned := make([]*Individual, 0)
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
	population := populationGenerator(&domain, 1)
	fitness(&domain, population) // TODO Change to useful individuals

	// Run ga
	stoppingCondition := true
	for i := 0; i < domain.Config.Generations && stoppingCondition; i++ {
		crossOverAmount := (domain.Config.Size * int(domain.Config.CrossOverRate))
		mutateAmount := (domain.Config.Size * int(domain.Config.MutationRate))
		carryAmount := domain.Config.Size - crossOverAmount - mutateAmount // TODO: Find out Anna if is guicci

		// evolve
		individualsOffspring := crossover(&domain, selection(&domain, population, crossOverAmount), selection, 2)
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
