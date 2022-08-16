package ga

func StubMutate(domain *Domain, individuals Individuals) Individuals {
	return individuals.ClonePopulation()
}
