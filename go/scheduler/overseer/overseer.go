package overseer

import (
	"context"
	"lib/logger"
	"lib/testutils"
	"lib/utils"
	"scheduler/data"
	"scheduler/ga"
	"time"
)

func WeeklyOverseer(schedulerData data.SchedulerData, schedulerConfig *data.SchedulerConfig) []data.Bookings {
	// Perform Magic
	var bookings []data.Bookings

	// Set default configurations
	var config data.Config
	config.Seed = 2
	config.PopulationSize = 15
	config.Generations = 1000
	config.MutationRate = 1.0
	config.CrossOverRate = 0.0
	config.TournamentSize = 5
	config.TimeLimitSeconds = 3

	if schedulerConfig != nil {
		if schedulerConfig.WeeklyConfig != nil {
			conf := schedulerConfig.WeeklyConfig
			// TODP: fix weekly crossover before allowing
			config.Seed = *utils.ReturnAltIfNil(&conf.Seed, &config.Seed)
			config.PopulationSize = *utils.ReturnAltIfNil(&conf.PopulationSize, &config.PopulationSize)
			config.Generations = *utils.ReturnAltIfNil(&conf.Generations, &config.Generations)
			config.CrossOverRate = *utils.ReturnAltIfNil(&conf.CrossOverRate, &config.CrossOverRate)
			config.MutationRate = *utils.ReturnAltIfNil(&conf.MutationRate, &config.MutationRate)
			config.TournamentSize = *utils.ReturnAltIfNil(&conf.TournamentSize, &config.TournamentSize)
			config.TimeLimitSeconds = *utils.ReturnAltIfNil(&conf.TimeLimitSeconds, &config.TimeLimitSeconds)
		}
	}

	// Create domain
	var domain ga.Domain
	domain.Terminals = data.ExtractUserIdsDuplicates(&schedulerData)
	domain.Config = &config
	domain.SchedulerData = &schedulerData

	// Create channel
	var c chan ga.Individual = make(chan ga.Individual)
	s, stopGA := context.WithCancel(context.Background())
	defer stopGA()

	go ga.GA(
		domain,
		func(domain *ga.Domain, individuals ga.Individuals, selectionFunc ga.Selection, offspring int) ga.Individuals {
			return ga.CrossoverCaller(ga.WeeklyDayVResourceCrossover, domain, individuals, selectionFunc, offspring)
		},
		func(domain *ga.Domain, individuals ga.Individuals) []float64 {
			return ga.WeeklyDayVResourceFitnessCaller(domain, individuals, ga.WeeklyDayVResourceFitnessValid)
		},
		ga.WeeklyDayVResourceMutateSwapValid,
		ga.WeeklyTournamentSelectionFitness,
		ga.WeeklyDayVResourcePopulationGenerator,
		c,
		&s,
	)

	// Listen on channel for best individual for x seconds
	var best ga.Individual
	best.Fitness = -1
	count := -1
	improvements := 0
	timeoutChanel := time.After(time.Second * time.Duration(config.TimeLimitSeconds))
	for {
		select {
		case <-timeoutChanel:
			logger.Debug.Println(testutils.Scolour(testutils.RED, "DEADLINE EXCEDED"))
			// logger.Error.Println("\n", best)
			// Stop the GA
			stopGA()
			bookings = append(bookings, best.ConvertIndividualToWeeklyBookings(domain))
			return bookings
		case candidate, ok := <-c: // if ok is false close event happened
			if !ok {
				bookings = append(bookings, best.ConvertIndividualToWeeklyBookings(domain))
				// logger.Debug.Println("\n", best)
				// logger.Debug.Println(testutils.Scolourf(testutils.PURPLE, "SOLUTIONS RECIEVED: %v, %v", count, improvements))
				return bookings
			}
			count++
			if candidate.Fitness > best.Fitness {
				// logger.Debug.Println(testutils.Scolourf(testutils.PURPLE, "IMPROVEMENT RECIEVED: %v, %v", count, improvements))
				logger.Debug.Println(testutils.Scolourf(testutils.PURPLE, "IMPROVEMENT RECIEVED: %v, %v, %v", count, improvements, candidate.Fitness))
				improvements++
				best = candidate
				// logger.Error.Printf("BEST INDIVIDUAL \n %v", best)
			}
		}
	}
}

func DailyOverseer(schedulerData data.SchedulerData, schedulerConfig *data.SchedulerConfig) []data.Bookings {
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
	config.TimeLimitSeconds = 3

	if schedulerConfig != nil {
		if schedulerConfig.DailyConfig != nil {
			conf := schedulerConfig.DailyConfig
			// TODP: fix weekly crossover before allowing
			config.Seed = *utils.ReturnAltIfNil(&conf.Seed, &config.Seed)
			config.PopulationSize = *utils.ReturnAltIfNil(&conf.PopulationSize, &config.PopulationSize)
			config.Generations = *utils.ReturnAltIfNil(&conf.Generations, &config.Generations)
			config.CrossOverRate = *utils.ReturnAltIfNil(&conf.CrossOverRate, &config.CrossOverRate)
			config.MutationRate = *utils.ReturnAltIfNil(&conf.MutationRate, &config.MutationRate)
			config.TournamentSize = *utils.ReturnAltIfNil(&conf.TournamentSize, &config.TournamentSize)
			config.TimeLimitSeconds = *utils.ReturnAltIfNil(&conf.TimeLimitSeconds, &config.TimeLimitSeconds)
		}
	}

	// Create domain
	var domain ga.Domain
	domain.Terminals = data.ExtractResourceIds(&schedulerData)
	domain.Config = &config
	domain.SchedulerData = &schedulerData
	domain.Map = data.ExtractUserIdMap(&schedulerData)
	domain.InverseMap = data.ExtractInverseUserIdMap(domain.Map)

	// Create channel
	var c chan ga.Individual = make(chan ga.Individual)
	s, stopGA := context.WithCancel(context.Background())
	defer stopGA()

	go ga.GA(
		domain,
		func(domain *ga.Domain, individuals ga.Individuals, selectionFunc ga.Selection, offspring int) ga.Individuals {
			return ga.CrossoverCaller(ga.PartiallyMappedFlattenCrossoverValid, domain, individuals, selectionFunc, offspring)
		},
		ga.DailyFitness,
		ga.DailyMutateValid,
		ga.WeeklyTournamentSelectionFitness,
		ga.DailyPopulationGeneratorValid,
		c,
		&s,
	)
	// go ga.GA(domain, ga.WeeklyStubCrossOver, ga.WeeklyStubFitness, ga.DailyMutate, ga.WeeklyTournamentSelection, ga.DailyPopulationGenerator, c, &s)

	// Listen on channel for best individual for x seconds
	var best ga.Individual
	best.Fitness = -1
	count := -1
	improvements := 0
	timeoutChanel := time.After(time.Second * time.Duration(config.TimeLimitSeconds))
	for {
		select {
		case <-timeoutChanel:
			logger.Debug.Println(testutils.Scolour(testutils.RED, "DEADLINE EXCEDED"))
			// Stop the GA
			stopGA()
			bookings = append(bookings, best.ConvertIndividualToDailyBookings(domain))
			// logger.Debug.Println(len(bookings))
			// logger.Debug.Println(best)
			return bookings
		case candidate, ok := <-c: // if ok is false close event happened
			if !ok {
				bookings = append(bookings, best.ConvertIndividualToDailyBookings(domain))
				// logger.Debug.Println(best)
				logger.Debug.Println(testutils.Scolourf(testutils.PURPLE, "SOLUTIONS RECIEVED: %v, %v", count, improvements))
				return bookings
			}
			count++
			if candidate.Fitness > best.Fitness {
				logger.Debug.Println(testutils.Scolourf(testutils.PURPLE, "IMPROVEMENT RECIEVED: %v, %v, %v", count, improvements, candidate.Fitness))
				improvements++
				best = candidate
				// logger.Error.Printf("BEST INDIVIDUAL \n %v", best)
			}
		}
	}
}
