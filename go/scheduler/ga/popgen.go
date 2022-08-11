package ga

import (
	"math/rand"
	"scheduler/data"
)

func StubPopulationGenerator(domain Domain, schedulerData data.SchedulerData, popSize int) Individuals {
	// Using representation Empl1 --- Empl2 --- Empl3 ...
	//						Mon       Mon       Thu   ...
	var population Individuals
	for j := 0; j < popSize; j++ {
		genes := []string{}
		for i := 0; i < len(schedulerData.Users); i++ {
			genes = append(genes, domain.Terminals[j%len(domain.Terminals)])
		}
		population = append(population, &Individual{Gene: genes})
	}
	return population
}

func WeeklyPopulationGenerator(domain Domain, schedulerData data.SchedulerData, popSize int) Individuals {
	// Using representation Empl1 --- Empl2 --- Empl3 ...
	//						Mon       Mon       Thu   ...
	var population Individuals
	for j := 0; j < popSize; j++ {
		genes := []string{}
		for i := 0; i < len(schedulerData.Users); i++ {
			genes = append(genes, domain.Terminals[rand.Intn(len(domain.Terminals))])
		}
		population = append(population, &Individual{Gene: genes})
	}
	return population
}
