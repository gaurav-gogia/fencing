package main

import (
	"strconv"
)

func count_compress(input string) []rune {
	runarry := []rune(input)
	length := len(runarry)

	for i := 0; i < length; {

		// Count characters
		count := counter(runarry, i, length)

		// Append character & count
		runarry = append(runarry, runarry[i])
		if count > 1 {
			countStr := strconv.Itoa(count)
			runarry = append(runarry, []rune(countStr)...)
		}

		// Remove counted characters
		runarry = runarry[count:]
		length = length - count
	}

	return runarry
}

func counter(runarry []rune, i, length int) int {
	count := 1
	for j := 1; j < length; j++ {

		if runarry[i] != runarry[j] {
			break
		}

		count++
	}

	return count
}
