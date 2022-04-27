package main

func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	direction := array[1] - array[0]
	for i := 2; i < len(array); i++ {
		if direction == 0 {
			direction = array[i] - array[i-1]
			continue
		}
		if breaksDireciton(direction, array[i-1], array[i]) {
			return false
		}
	}
	return true
}

func breaksDireciton(direction, previousInt, currentInt int) bool {
	difference := currentInt - previousInt
	if direction > 0 {
		return difference < 0
	}
	return difference > 0
}
