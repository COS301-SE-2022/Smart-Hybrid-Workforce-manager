package ga

import (
	cu "lib/collectionutils"
	"lib/utils"
	"scheduler/data"
)

// Domain represents domain information to the problem
type Domain struct {
	// Weekly scheduler: User array with user id's duplicated for amount per week
	Terminals []string

	Config        *data.Config
	SchedulerData *data.SchedulerData
	Map           map[int](string)
	InverseMap    map[string]([]int) // Array due to the assumption Map may have a many-to-one relationship
	// If true, existing bookings may be rescheduled
	Reschedule bool
}

func (domain *Domain) GetRandomTerminal() string {
	return domain.Terminals[utils.RandInt(0, len(domain.Terminals))]
}

func (domain *Domain) GetRandomTerminalArrays(length int) []string {
	var result []string
	for i := 0; i < length; i++ {
		result = append(result, domain.GetRandomTerminal())
	}
	return result
}

// Gets unique* elements from the terminals array
// unique in this context means that it will not take the exact same element twice,
// however if duplicates are present, it could happen that an element gets selected twice
// if len(terminals) < length, panic will result
func (domain *Domain) GetRandomUniqueTerminalArrays(length int) []string {
	domainTerminalsCopy := cu.Copy1DArr(domain.Terminals)
	var result []string
	for i := 0; i < length; i++ {
		randi := utils.RandInt(0, len(domainTerminalsCopy))
		result = append(result, domainTerminalsCopy[randi])
		domainTerminalsCopy = cu.RemElemenAtI(domainTerminalsCopy, randi)
	}
	return result
}

// TODO: @JonathanEnslin look at moving this piece of code into the domain as well since it is common across individuals.
// A map that contains the user indices per team.
// GetTeamUserIndices gets a map where the keys are teamIds, and the value is an array indicating which indices in a daily individual
// belongs to the tean
func (domain *Domain) GetTeamUserIndices() map[string][]int {
	userIndicesMap := domain.InverseMap
	teamInfos := domain.SchedulerData.Teams
	teamUserIndices := make(map[string][]int) // map[teamId][user indices]
	for _, team := range teamInfos {          // For each team
		teamUserIndices[*team.Id] = make([]int, 0)
		for _, userId := range team.UserIds {
			// Add the gene index of the user to the team indices array, to indicate which indices in the gene
			// map to users part of the team
			teamUserIndices[*team.Id] = append(teamUserIndices[*team.Id], userIndicesMap[userId]...)
		}
	}
	return teamUserIndices
}
