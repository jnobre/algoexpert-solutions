package main

func SelectionSort(array []int) []int {

	currentIndex := 0
	for currentIndex < len(array)-1 {
		smallestIndex := currentIndex

		for i := currentIndex + 1; i < len(array); i++ {
			if array[smallestIndex] > array[i] {
				smallestIndex = i
			}
		}
		array[smallestIndex], array[currentIndex] = array[currentIndex], array[smallestIndex]

		currentIndex++
	}
	return array
}
