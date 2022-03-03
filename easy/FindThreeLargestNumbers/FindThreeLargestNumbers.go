package main

import "math"

func FindThreeLargestNumbers(array []int) []int {
	three := []int{math.MinInt64, math.MinInt64, math.MinInt64}
	for _, num := range array {
		if num > three[2] {
			shiftAndUpdate(three, num, 2)
		} else if num > three[1] {
			shiftAndUpdate(three, num, 1)
		} else if num > three[0] {
			shiftAndUpdate(three, num, 0)
		}
	}
	return three
}

func shiftAndUpdate(three []int, num, idx int) {
	for i := 0; i < idx+1; i++ {
		if i == idx {
			three[i] = num
		} else {
			three[i] = three[i+1]
		}
	}
}
