package ga

// func StubPopulationGenerator(domain Domain, schedulerData data.SchedulerData, popSize int) Individuals {
// 	// Using representation Empl1 --- Empl2 --- Empl3 ...
// 	//						Mon       Mon       Thu   ...
// 	var population Individuals
// 	for j := 0; j < popSize; j++ {
// 		genes := []string{}
// 		for i := 0; i < len(schedulerData.Users); i++ {
// 			genes = append(genes, domain.Terminals[j%len(domain.Terminals)])
// 		}
// 		population = append(population, &Individual{Gene: genes})
// 	}
// 	return population
// }

func populationGenerator(domain *Domain, popSize int) Individuals {
	// Using representation 1 (DayVResource)
	schedulerData := domain.SchedulerData

	// Get slot sizes/day
	slotSizes := []int{}
	for i := 0; i < 7; i++ {
		slotSizes = append(slotSizes, 0)
	}
	for _, booking := range *schedulerData.CurrentBookings. {
		var day int = int(booking.bo)
	}

}

// func WeeklyPopulationGenerator(domain Domain, schedulerData data.SchedulerData, popSize int) Individuals {
// 	// Using representation Empl1 --- Empl2 --- Empl3 --- Empl1
// 	//						Mon       Mon       Thu       Tue
// 	//						Tue
// 	var population Individuals
// 	for j := 0; j < popSize; j++ {
// 		genes := []string{}
// 		for i := 0; i < len(schedulerData.Users); i++ {
// 			genes = append(genes, domain.Terminals[rand.Intn(len(domain.Terminals))])
// 		}
// 		population = append(population, &Individual{Gene: genes})
// 	}
// 	return population
// }
