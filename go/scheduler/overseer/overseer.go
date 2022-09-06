package overseer

import (
	"fmt"
	"scheduler/data"
	"scheduler/ga"
)

func WeeklyOverseer(schedulerData data.SchedulerData) []data.Bookings {
	// Perform Magic
	var bookings []data.Bookings

	// Set configurations
	var config data.Config
	config.Seed = 2
	config.PopulationSize = 150
	config.Generations = 100
	config.MutationRate = 0.25
	config.CrossOverRate = 0.65
	config.TournamentSize = 15

	// Create domain
	var domain ga.Domain
	domain.Terminals = data.ExtractUserIdsDuplicates(&schedulerData)
	domain.Config = &config
	domain.SchedulerData = &schedulerData

	results := ga.GA(domain, ga.WeeklyDayVResourceCrossover, ga.WeeklyDayVResourceFitness, ga.WeeklyDayVResourceMutateSwapValid, ga.WeeklyTournamentSelection, ga.WeeklyDayVResourcePopulationGenerator)
	// results := ga.GA(domain, ga.WeeklyDayVResourceCrossover, ga.WeeklyDayVResourceFitness, ga.WeeklyDayVResouceMutate, ga.WeeklyTournamentSelection, ga.WeeklyDayVResourcePopulationGenerator)

	if len(results) == 0 { // todo add check

	}

	// Get best individual
	for i, indiv := range results {
		// todo put through validation function

		// transform into what the backend needs
		if i == 0 {
			bookings = append(bookings, indiv.ConvertIndividualToWeeklyBookings(domain))
		}
	}

	return bookings
}

func DailyOverseer(schedulerData data.SchedulerData) []data.Bookings {
	// Perform Magic
	var bookings []data.Bookings

	// Set configurations
	var config data.Config
	config.Seed = 2
	config.PopulationSize = 150
	config.Generations = 100
	config.MutationRate = 0.45
	config.CrossOverRate = 0.45
	config.TournamentSize = 10

	// Create domain
	var domain ga.Domain
	domain.Terminals = data.ExtractResourceIds(&schedulerData)
	domain.Config = &config
	domain.SchedulerData = &schedulerData
	domain.Map = data.ExtractUserIdMap(&schedulerData)

	indvs := ga.DailyPopulationGenerator(&domain, 1)

	for _, indiv := range indvs {
		fmt.Println(indiv.StringDomain(domain))
		// mutated := ga.DailyMutate(&domain, ga.Individuals{indiv})
		// fmt.Println(mutated[0].StringDomain(domain))
		bookings = append(bookings, indiv.ConvertIndividualToDailyBookings(domain))
	}

	return bookings
}
