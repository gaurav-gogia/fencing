package main

func substring_pallindrome(input string) string {
	if isPallindrome(input) {
		return input
	}

	outs := getAllSubStrings(input)
	pallmap := make(map[int]string)

	for _, out := range outs {
		if isPallindrome(out) {
			pallmap[len(out)] = out
		}
	}

	return getmax(pallmap)
}

func isPallindrome(in string) bool {
	return in == reverse(in)
}

func reverse(input string) string {
	runes := []rune(input)
	length := len(runes) - 1

	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func getAllSubStrings(input string) []string {
	var subStrings []string
	for i := range input {

		for length := i; length <= len(input); length++ {

			sub := input[i:length]
			if len(sub) > 1 {
				subStrings = append(subStrings, sub)
			}

		}

	}

	return subStrings
}

func getmax(pallmap map[int]string) string {
	var max int
	for k, _ := range pallmap {
		if max < k {
			max = k
		}
	}
	return pallmap[max]
}
