package main

func FirstNonRepeatingCharacter(str string) int {
	characterCounts := map[string]int{}

	for _, char := range str {
		characterCounts[string(char)] += 1
	}

	for i, char := range str {
		if characterCounts[string(char)] == 1 {
			return i
		}
	}
	return -1
}
