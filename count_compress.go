package main

import (
	"strconv"
)

func count_compress(input string) []rune {
	var count, i, j int
	runarry := []rune(input)
	length := len(runarry)

	for i = 0; i < length; i++ {
		count = 1

		for j = i + 1; j < length; j++ {

			if runarry[i] != runarry[j] {
				break
			}

			count++
		}

		runarry = append(runarry, runarry[i])
		if count > 1 {
			countStr := strconv.Itoa(count)
			runarry = append(runarry, []rune(countStr)...)
		}

		i = j - 1
	}

	return runarry[i:]
}
