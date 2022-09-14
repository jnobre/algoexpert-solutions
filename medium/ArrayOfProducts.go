package main

func ArrayOfProducts(array []int) []int {
	products := make([]int, len(array))
	for i := range array {
		products[i] = 1
	}

	leftRunningProduct := 1
	for i := range array {
		products[i] = leftRunningProduct
		leftRunningProduct *= array[i]
	}

	rightRunningProduct := 1
	for i := len(array) - 1; i >= 0; i-- {
		products[i] *= rightRunningProduct
		rightRunningProduct *= array[i]
	}
	return products
}
