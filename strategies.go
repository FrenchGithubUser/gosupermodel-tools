package main

//
// each strategy tells which group to remove from a given situation
//

// this strategy chooses the biggest group, and if there are multiple biggest ones, the one that is the most on the bottom right is choosen
func strategyBiggestFirst(groups *map[string][][][2]int) ([][2]int, string) {
	// Find the largest group on the grid
	largestGroup := [][2]int{}
	largestGroupName := ""
	largestSize := 0
	for itemName, groups := range *groups {
		for _, group := range groups {
			if len(group) > largestSize {
				largestSize = len(group)
				largestGroup = group
				largestGroupName = itemName
			}
		}
	}

	return largestGroup, largestGroupName
}
