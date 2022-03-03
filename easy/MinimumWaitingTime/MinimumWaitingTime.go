package main

import "sort"

func MinimumWaitingTime(queries []int) int {
	// Write your code here.
	sort.Ints(queries)

	var total int
	for idx, duration := range queries {
		queriesLeft := len(queries) - (idx + 1)
		total += queriesLeft * duration
	}

	return total
}
