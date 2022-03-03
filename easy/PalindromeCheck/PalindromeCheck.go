package main

func IsPalindrome(str string) bool {
	var middle = -1
	if len(str)%2 == 1 {
		middle = len(str) / 2
	}

	var runeValues []rune
	for i, runeValue := range str {
		if i == middle {
			continue
		}

		if len(runeValues) > 0 && runeValues[len(runeValues)-1] == runeValue {
			_, runeValues = runeValues[len(runeValues)-1], runeValues[:len(runeValues)-1]
		} else {
			runeValues = append(runeValues, runeValue)
		}
	}

	if len(runeValues) == 0 {
		return true
	}

	return false
}
