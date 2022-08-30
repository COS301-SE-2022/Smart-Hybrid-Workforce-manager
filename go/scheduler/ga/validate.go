package ga

import "lib/utils"

// Makes an individual valid
func ValidateIndividual(domain *Domain, indiv *Individual) {
	// First remove duplicates on a single day, at the same time, build maps
	usersComingInOnDay := make([]map[string]int, len(indiv.Gene)) // map[user id] (index of their slot)
	daysThatUsersComeIn := make(map[string][]int)                 // map[user id] {int array, where int corresponds to the day a user comes in}
	for i := 0; i < len(indiv.Gene); i++ {
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
			}
		}
	}
	// Count how many times a week a user needs to come in
	numTimesToComeInPerUser := make(map[string]int) // map[user id]#times need to come in
	for _, uid := range domain.Terminals {
		numTimesToComeInPerUser[uid]++
	}
	// Perform validation to check how many times a week they have to come in
	for userid, numTimesToComeIn := range numTimesToComeInPerUser {
		if numTimesToComeIn < len(daysThatUsersComeIn[userid]) { // If user is scheduled to come in too much
			// remove randomly
			for i := 0; i < len(daysThatUsersComeIn[userid])-numTimesToComeIn; i++ {
				// Get a randomIndex to select a day from
				randi := utils.RandInt(0, len(daysThatUsersComeIn[userid]))
				randDay := daysThatUsersComeIn[userid][randi]
				// Remove user from that day
				for sloti, uid := range indiv.Gene[randDay] {
					if uid == userid {
						indiv.Gene[randDay][sloti] = ""
						break
					}
				}
			}
		}
	}
}
