package ga

func StubFitness(domain *Domain, individuals Individuals) []float64 {
	var result []float64
	for i := 0; i < len(individuals); i++ {
		result = append(result, 0.0)
		individuals[i].Fitness = 0.0
	}
	return result
}

func DayVResourceFitness(domain *Domain, individuals Individuals) []float64 {
	var result []float64
	for _, individual := range individuals {
		result = append(result, dayVResourceFitness(domain, individual))
	}
	return result
}

func dayVResourceFitness(domain *Domain, individual *Individual) float64 {
	differentUsersCount := individual.DifferentUsersCount(domain)
	doubleBookings := individual.DoubleBookingsCount(domain)
	usersNotCommingInTheirSpecifiedAmountCount := individual.UsersNotCommingInTheirSpecifiedAmountCount(domain)
	teamsAttendingSameDays := individual.TeamsAttendingSameDays(domain)
	return float64(differentUsersCount) * teamsAttendingSameDays / (float64(doubleBookings) * float64(usersNotCommingInTheirSpecifiedAmountCount))
}

func (individual *Individual) DifferentUsersCount(domain *Domain) int {
	return 0
}

func (individual *Individual) DoubleBookingsCount(domain *Domain) int {
	return 1
}

func (individual *Individual) UsersNotCommingInTheirSpecifiedAmountCount(domain *Domain) int {
	return 1
}

func (individual *Individual) TeamsAttendingSameDays(domain *Domain) float64 {
	return 0.0
}
