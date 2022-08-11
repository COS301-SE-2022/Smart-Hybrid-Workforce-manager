package ga

import (
	"lib/utils"
	"math/rand"
	"scheduler/data"
)

type Crossover func(domain Domain, individuals Individuals, selectionFunc Selection, fitnii []float64, offspring int) Individuals
type Fitness func(domain Domain, individuals Individuals) []float64
type Mutate func(domain Domain, individuals Individuals) Individuals
type Selection func(domain Domain, individuals Individuals, fitness []float64, count int) Individuals
type PopulationGenerator func(domain Domain, schedulerData data.SchedulerData, popSize int) Individuals

type Domain struct {
	Terminals []string
}

func (domain *Domain) GetRandomTerminal() string {
	return domain.Terminals[utils.RandInt(0, len(domain.Terminals))]
}

type Individual struct {
	Gene    []string
	Fitness int
}

type Individuals []*Individual

func (individual *Individual) Clone() *Individual {
	newIndividual := &Individual{Gene: individual.Gene}
	return newIndividual
}

func (population Individuals) ClonePopulation() []*Individual {
	cloned := make([]*Individual, 0)
	for _, individual := range population {
		cloned = append(cloned, individual.Clone())
	}
	return cloned
}

func GA(schedulerData data.SchedulerData, config data.Config, domain Domain, crossover Crossover, fitness Fitness, mutate Mutate, selection Selection, populationGenerator PopulationGenerator) Individuals {
	// Seed
	rand.Seed(int64(config.Seed))

	// Create initial pop and calculate fitnesses
	population := populationGenerator(domain, schedulerData, 1)
	fitnii := fitness(domain, population) // TODO Change to useful individuals

	// Run ga
	stoppingCondition := true
	for i := 0; i < config.Generations && stoppingCondition; i++ {
		crossOverAmount := (config.Size * int(config.CrossOverRate))
		mutateAmount := (config.Size * int(config.MutationRate))
		carryAmount := config.Size - crossOverAmount - mutateAmount // TODO: Find out Anna if is guicci
		// evolve
		individualsOffspring := crossover(domain, selection(domain, population, fitnii, crossOverAmount), selection, fitnii, 2)
		individualsMutated := mutate(domain, selection(domain, population, fitnii, mutateAmount))
		individualsCarry := selection(domain, population, fitnii, carryAmount)

		population := append(individualsOffspring, individualsMutated...)
		population = append(population, individualsCarry...)

		fitnii = fitness(domain, population)
	}

	return nil
}

//       Monday   -   Tuesday   -  Wednesday
// 08:00 emp1, emp2
// 09:00
// 10:10

// emp1     emp2    emp3
// Mon		Mon
// 0-1		0-1
