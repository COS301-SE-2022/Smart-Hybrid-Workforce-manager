package ga

import (
	"lib/collectionutils"
	"lib/utils"
)

///////////////////////////////////////////////////
// WEEKLY

// Makes an individual valid
func ValidateIndividual(domain *Domain, indiv *Individual) {
	// logger.Error.Printf("len of domain terminals: %v   -   %v", len(domain.Terminals), domain.Terminals)
	// First remove duplicates on a single day, at the same time, build maps
	usersComingInOnDay := make([]map[string]int, len(indiv.Gene)) // map[user id] (index of their slot)
	daysThatUsersComeIn := make(map[string][]int)                 // map[user id] {int array, where int corresponds to the day a user comes in}
	openSlots := make([][]int, len(indiv.Gene))                   // Used to keep track of open spots throughout the individual
	for i := 0; i < len(indiv.Gene); i++ {
		openSlots[i] = []int{}
		usersComingInOnDay[i] = make(map[string]int)
		for sloti, userid := range indiv.Gene[i] {
			_, userIdAlreadyIn := usersComingInOnDay[i][userid]
			if userid != "" && !userIdAlreadyIn { // only if user has not already been mapped
				if _, ok := daysThatUsersComeIn[userid]; !ok {
					daysThatUsersComeIn[userid] = []int{}
				}
				daysThatUsersComeIn[userid] = append(daysThatUsersComeIn[userid], i)
				usersComingInOnDay[i][userid] = sloti // false here has no meaning
			} else if userIdAlreadyIn {
				indiv.Gene[i][sloti] = "" // Remove duplicate user
			} else if userid == "" {
				openSlots[i] = append(openSlots[i], sloti)
			}
		}
	}
	// Count how many times a week a user needs to come in
	numTimesToComeInPerUser := make(map[string]int) // map[user id]#times need to come in
	for _, uid := range domain.Terminals {
		numTimesToComeInPerUser[uid]++
	}
	// logger.Error.Printf(" NUM TIMES TO COME IN: %v", numTimesToComeInPerUser)
	// Perform validation to check how many times a week they have to come in
	for userid, numTimesToComeIn := range numTimesToComeInPerUser {
		if numTimesToComeIn < len(daysThatUsersComeIn[userid]) { // If user is scheduled to come in too much
			// fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAA")
			// remove randomly
			for i := 0; i < len(daysThatUsersComeIn[userid])-numTimesToComeIn; i++ {
				// Get a randomIndex to select a day from
				randi := utils.RandInt(0, len(daysThatUsersComeIn[userid]))
				randDay := daysThatUsersComeIn[userid][randi]
				// Remove user from that day
				for sloti, uid := range indiv.Gene[randDay] {
					if uid == userid {
						indiv.Gene[randDay][sloti] = ""
						openSlots[randDay] = append(openSlots[randDay], sloti)
						break
					}
				}
			}
		} else if numTimesToComeIn > len(daysThatUsersComeIn[userid]) { // user not comingin enough times
			for i := 0; i < numTimesToComeIn-len(daysThatUsersComeIn[userid]); i++ {
				// fmt.Println("JSHDKAJSHDAJSHDKJHAKSDJHSKDJH")
				// get random open slot
				randDay := utils.RandInt(0, len(openSlots))
				// choose day with most open slots
				if len(openSlots[randDay]) == 0 {
					maxOpen := len(openSlots[0])
					maxIndex := 0
					for i := 1; i < len(openSlots); i++ {
						if len(openSlots[i]) > maxOpen {
							maxOpen = len(openSlots[i])
							maxIndex = i
						}
					}
					randDay = maxIndex
				}
				if len(openSlots[randDay]) <= 0 {
					continue // no space to fix
				}
				randSloti := utils.RandInt(0, len(openSlots[randDay]))
				// add user to select slot
				indiv.Gene[randDay][randSloti] = userid
				// remove random slot
				openSlots[randDay] = collectionutils.RemElemenAtI(openSlots[randDay], randSloti)
			}
		}
	}
}

///////////////////////////////////////////////////
// DAILY

func (individual *Individual) CheckIfValidDaily() bool {
	gene := individual.Gene
	// Will be used to check if any resources are assigned twice
	_, resourceIdGroups := collectionutils.GroupBy(gene[0], func(id string) string { return id })
	for _, group := range resourceIdGroups {
		if len(group) > 1 {
			return false
		}
	}
	return true
}

func (individual *Individual) ValidateIndividual(domain *Domain) {
	gene := individual.Gene
	// Will be used to check if any resources are assigned twice
	_, resourceIdGroups := collectionutils.GroupBy(gene[0], func(id string) string { return id })
	availableTerminals := collectionutils.SliceDifference(domain.Terminals, gene[0])
	if len(availableTerminals) == 0 {
		return
	}
	for i, id := range gene[0] {
		if len(resourceIdGroups[id]) > 1 {
			randi := utils.RandInt(0, len(availableTerminals))
			newId := availableTerminals[randi]
			availableTerminals = collectionutils.RemElemenAtI(availableTerminals, randi)
			resourceIdGroups[id] = collectionutils.RemElemenAtI(resourceIdGroups[id], 0)
			gene[0][i] = newId
		}
	}
}
