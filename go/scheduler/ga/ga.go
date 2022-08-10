package ga

import "scheduler/data"

type Crossover func(individuals []*Individual, offspring int) []*Individual
type Fitness func(individuals []*Individual) []float64
type Mutate func(individuals []*Individual) []*Individual
type Selection func(individuals []*Individual, fitness []float64, count int) []*Individual
type PopulationGenerator func(schedulerData data.SchedulerData) []*Individual

type Individual struct {
}

func GA(schedulerData data.SchedulerData, config data.Config, crossover Crossover, fitness Fitness, mutate Mutate, selection Selection, populationGenerator PopulationGenerator) {
	// Create initial pop and calculate fitnesses
	population := populationGenerator(schedulerData)
	fitnii := fitness(population) // TODO Change to useful individuals

	// Run ga
	stoppingCondition := true
	for i := 0; i < config.Generations && stoppingCondition; i++ {
		crossOverAmount := (config.Size * int(config.CrossOverRate))
		mutateAmount := (config.Size * int(config.MutationRate))
		carryAmount := config.Size - crossOverAmount - mutateAmount // TODO: Find out Anna if is guicci

		// evolve
		individualsOffspring := crossover(selection(population, fitnii, crossOverAmount), 2)
		individualsMutated := mutate(selection(population, fitnii, mutateAmount))
		individualsCarry := selection(population, fitnii, carryAmount)

		population := append(individualsOffspring, individualsMutated...)
		population = append(population, individualsCarry...)

		fitnii = fitness(population)
	}
}

//       Monday   -   Tuesday   -  Wednesday
// 08:00 emp1, emp2
// 09:00
// 10:10

// emp1     emp2    emp3
// Mon		Mon
// 0-1		0-1
