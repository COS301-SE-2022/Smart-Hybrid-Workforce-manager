package ga

import (
	"context"
	"fmt"
	"lib/logger"
	"lib/testutils"
	"lib/utils"
	"math"
	"math/rand"
	"scheduler/data"
	"strings"
	"time"
)

// Individuals is a collection of solutions to the domain problem
type Individuals []*Individual

type Crossover func(domain *Domain, individuals Individuals, selectionFunc Selection, offspring int) Individuals
type Fitness func(domain *Domain, individuals Individuals) []float64
type Mutate func(domain *Domain, individuals Individuals) Individuals
type Selection func(domain *Domain, individuals Individuals, count int) Individuals
type PopulationGenerator func(domain *Domain, popSize int) Individuals

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

// ConvertIndividualToWeeklyBookings converts an individual result to bookings
func (individual *Individual) ConvertIndividualToWeeklyBookings(domain Domain) data.Bookings {
	var bookings data.Bookings

	if individual == nil {
		return nil
	}

	for i, day := range individual.Gene {
		for _, user := range day {
			if user == "" {
				continue
			}
			for _, userInfo := range domain.SchedulerData.Users {
				if *userInfo.Id == user {
					booked := true
					start := time.Date(domain.SchedulerData.StartDate.Year(), domain.SchedulerData.StartDate.Month(), domain.SchedulerData.StartDate.Day(), userInfo.PreferredStartTime.Hour(), userInfo.PreferredStartTime.Minute(), 0, 0, time.UTC)
					end := time.Date(domain.SchedulerData.StartDate.Year(), domain.SchedulerData.StartDate.Month(), domain.SchedulerData.StartDate.Day(), userInfo.PreferredEndTime.Hour(), userInfo.PreferredEndTime.Minute(), 0, 0, time.UTC)
					start = start.AddDate(0, 0, i)
					end = end.AddDate(0, 0, i)
					userId := user
					desk := "DESK"
					bookings = append(bookings, &data.Booking{UserId: &userId, Start: &start, End: &end, Booked: &booked, Automated: &booked, ResourceType: &desk})
				}
			}
		}
	}

	return bookings
}

// ConvertIndividualToDailyBookings converts an individual result to bookings
func (individual *Individual) ConvertIndividualToDailyBookings(domain Domain) data.Bookings {
	var bookings data.Bookings

	if individual == nil {
		return nil
	}

	for i, booking := range *domain.SchedulerData.CurrentBookings {
		booking.ResourceId = &individual.Gene[0][i]
		bookings = append(bookings, booking)
	}

	return bookings
}

// GetRandomIndividual retrieves a random individual in the population
func (population Individuals) GetRandomIndividual() *Individual {
	return population[utils.RandInt(0, len(population))]
}

func calculateGAStats(population Individuals) (maxi int, avg, maxFitness, minFitness float64) {
	// Get max and avg fitness
	totalFitness := 0.0
	maxFitness = math.Inf(-1)
	minFitness = math.Inf(1)
	for i, indiv := range population {
		maxFitness = math.Max(maxFitness, indiv.Fitness)
		minFitness = math.Min(minFitness, indiv.Fitness)
		if indiv.Fitness > maxFitness {
			maxi = i
		}
		totalFitness += indiv.Fitness
	}
	avg = totalFitness / float64(len(population))
	return
}

func printGAGraphs(multiplier, maxMultiplier, avg, maxFitness, minFitness float64) {
	fmt.Printf("METRICS: MAX=%f   MIN=%f  AVG=%f\n", maxFitness, minFitness, avg)
	if int(multiplier*avg) <= 0 {
		logger.Debug.Print("-")
	} else {
		logger.Debug.Print(strings.Repeat("*", int(multiplier*avg)))
	}
	if int(maxMultiplier*maxFitness) <= 0 {
		logger.Debug.Print("-")
	} else {
		logger.Debug.Print(strings.Repeat("=", int(multiplier*maxFitness)))
	}
}

func (population Individuals) getBestI() int {
	maxi := 0
	maxFitness := population[maxi].Fitness
	for indivi, indiv := range population {
		if indiv.Fitness > maxFitness {
			maxi = indivi
		}
		maxFitness = math.Max(maxFitness, indiv.Fitness)
	}
	return maxi
}

// GA is a generic configurable genetic algorithm that produces multiple solutions to the domain problem
func GA(domain Domain, crossover Crossover, fitness Fitness, mutate Mutate, selection Selection, populationGenerator PopulationGenerator, solutionChannel chan Individual, forceStop *context.Context) {
	// Seed
	rand.Seed(int64(domain.Config.Seed))
	//start := time.Now()

	// Create initial pop and calculate fitnesses
	population := populationGenerator(&domain, domain.Config.PopulationSize)

	fitness(&domain, population)

	_, avg, maxFitness, minFitness := calculateGAStats(population)
	// multiplier := 50.0 / avg
	maxMultiplier := 40.0 / maxFitness
	printGAGraphs(maxMultiplier, maxMultiplier, avg, maxFitness, minFitness)

	// return selection(&domain, population, 1)
	logger.Debug.Println("\n", *selection(&domain, population, 1)[0])
	// Run ga
	stoppingCondition := true
	for i := 0; i < domain.Config.Generations && stoppingCondition; i++ {
		if i%10 == 0 {
			logger.Debug.Println(testutils.Scolourf(testutils.BLUE, "Generation %v", i))
		}

		// Check if GA must be stopped after this run
		stoppingCondition = (*forceStop).Err() == nil

		// Calculate the amounts
		crossOverAmount := int(float64(domain.Config.PopulationSize) * domain.Config.CrossOverRate)
		mutateAmount := int(float64(domain.Config.PopulationSize) * domain.Config.MutationRate)
		carryAmount := domain.Config.PopulationSize - crossOverAmount - mutateAmount // TODO: Find out Anna if is guicci

		// evolve
		individualsOffspring := crossover(&domain, population, selection, crossOverAmount)

		// Validate crossover individuals
		for i := 0; i < len(individualsOffspring); i++ {
			ValidateIndividual(&domain, individualsOffspring[i]) // Maybe works, maybs not
		}

		// Get mutated individuals
		individualsMutated := mutate(&domain, selection(&domain, population, mutateAmount))

		// Get carried individuals
		individualsCarry := selection(&domain, population, carryAmount)

		// Create new population
		population = append(individualsOffspring, individualsMutated...)
		population = append(population, individualsCarry...)

		// Calculate the fitness of the population
		fitness(&domain, population)

		if i%10 == 0 {
			_, avg, maxFitness, minFitness = calculateGAStats(population)
			printGAGraphs(maxMultiplier, maxMultiplier, avg, maxFitness, minFitness)
		}

		maxi := population.getBestI()
		// send individual on channel
		solutionChannel <- *population[maxi]
	}
	_, avg, maxFitness, minFitness = calculateGAStats(population)
	printGAGraphs(maxMultiplier, maxMultiplier, avg, maxFitness, minFitness)

	// Send final best individual
	maxi := population.getBestI()
	solutionChannel <- *population[maxi]

	close(solutionChannel)
	logger.Debug.Println("\n", *population[maxi], "\n", maxFitness)
}

// String method for printing individuals
func (individual Individual) String() string {
	// Returns table representation of an individual
	if len(individual.Gene) == 0 {
		return "no elements"
	}
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
		if len(individual.Gene) == 1 {
			table[i][0] = fmt.Sprintf("%38s|", "Resources")
		} else {
			table[i][0] = fmt.Sprintf("%38s|", fmt.Sprintf("Day-%d", i))
		}
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

// String function for a string with a map
func (individual *Individual) StringDomain(domain Domain) string {
	// Returns table representation of an individual
	if len(individual.Gene) == 0 || len(individual.Gene[0]) == 0 {
		return "no elements"
	}

	maxSlotsize := -1
	for _, day := range individual.Gene {
		if len(day) > maxSlotsize {
			maxSlotsize = len(day)
		}
	}
	table := make([][]string, len(individual.Gene)+1) // +1 for user_id column
	for i := range table {
		table[i] = make([]string, maxSlotsize+3) // +1 for day thing + 1 for ending border thingy
	}

	for i := 0; i < len(individual.Gene); i++ {
		for j := 0; j < len(table[i]); j++ {
			table[i][j] = strings.Repeat(" ", 39)
		}
	}

	for i := range individual.Gene {
		if len(individual.Gene) == 1 {
			table[i][0] = fmt.Sprintf("%38s|", "Resources")
		} else {
			table[i][0] = fmt.Sprintf("%38s|", fmt.Sprintf("Day-%d", i))
		}
		table[i][1] = strings.Repeat("=", 39)
	}
	table[1][0] = fmt.Sprintf("%38s|", "User_ids") // user id column added
	table[1][1] = strings.Repeat("=", 39)

	for i := 0; i < len(individual.Gene)+1; i++ {
		counter := 0
		if i < len(individual.Gene) {
			for j := 0; j < len(individual.Gene[i]); j++ {
				counter = j
				table[i][j+2] = individual.Gene[i][j]
				table[i][j+2] = fmt.Sprintf("%38s|", individual.Gene[i][j])
			}
		} else {
			for j := 0; j < len(individual.Gene[i-1]); j++ {
				counter = j
				table[i][j+2] = fmt.Sprintf("%38s|", domain.Map[j])
			}
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
