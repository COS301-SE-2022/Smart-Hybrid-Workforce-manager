package overseer

import (
	"context"
	"lib/logger"
	"lib/testutils"
	"scheduler/data"
	"scheduler/ga"
	"time"
)

func WeeklyOverseer(schedulerData data.SchedulerData) []data.Bookings {
	// Perform Magic
	var bookings []data.Bookings

	// Set configurations
	var config data.Config
	config.Seed = 2
	config.PopulationSize = 15
	config.Generations = 100
	config.MutationRate = 0.25
	config.CrossOverRate = 0.65
	config.TournamentSize = 3

	// Create domain
	var domain ga.Domain
	domain.Terminals = data.ExtractUserIdsDuplicates(&schedulerData)
	domain.Config = &config
	domain.SchedulerData = &schedulerData

	// Create channel
	var c chan ga.Individual = make(chan ga.Individual)
	var s context.Context = context.Background()

	go ga.GA(domain, ga.WeeklyDayVResourceCrossover, ga.WeeklyDayVResourceFitness, ga.WeeklyDayVResourceMutateSwapValid, ga.WeeklyTournamentSelection, ga.WeeklyDayVResourcePopulationGenerator, c, &s)

	// Listen on channel for best individual for x seconds
	var best ga.Individual
	best.Fitness = -1
	for {
		select {
		case <-time.After(time.Second * 2):
			logger.Debug.Println(testutils.Scolour(testutils.RED, "DEADLINE EXCEDED"))
			s.Done()
			bookings = append(bookings, best.ConvertIndividualToWeeklyBookings(domain))
			return bookings
		case candidate, ok := <-c: // if ok is false close event happened
			if !ok {
				bookings = append(bookings, best.ConvertIndividualToWeeklyBookings(domain))
				return bookings
			}
			if candidate.Fitness > best.Fitness {
				best = candidate
				// logger.Error.Printf("BEST INDIVIDUAL \n %v", best)
			}
		}
	}
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
	domain.InverseMap = data.ExtractInverseUserIdMap(domain.Map)

	// Create channel
	var c chan ga.Individual = make(chan ga.Individual)
	var s context.Context = context.Background()

	go ga.GA(
		domain,
		func(domain *ga.Domain, individuals ga.Individuals, selectionFunc ga.Selection, offspring int) ga.Individuals {
			return ga.CrossoverCaller(ga.PartiallyMappedFlattenCrossoverValid, domain, individuals, selectionFunc, offspring)
		},
		ga.DailyFitness,
		ga.DailyMutateValid,
		ga.WeeklyTournamentSelection,
		ga.DailyPopulationGeneratorValid,
		c,
		&s,
	)
	// go ga.GA(domain, ga.WeeklyStubCrossOver, ga.WeeklyStubFitness, ga.DailyMutate, ga.WeeklyTournamentSelection, ga.DailyPopulationGenerator, c, &s)

	// Listen on channel for best individual for x seconds
	var best ga.Individual
	best.Fitness = -1
	for {
		select {
		case <-time.After(time.Second * 1):
			logger.Debug.Println(testutils.Scolour(testutils.RED, "DEADLINE EXCEDED"))
			s.Done()
			bookings = append(bookings, best.ConvertIndividualToDailyBookings(domain))
			return bookings
		case candidate, ok := <-c: // if ok is false close event happened
			if !ok {
				bookings = append(bookings, best.ConvertIndividualToDailyBookings(domain))
				return bookings
			}
			if candidate.Fitness > best.Fitness {
				best = candidate
				// logger.Error.Printf("BEST INDIVIDUAL \n %v", best)
			}
		}
	}
}
