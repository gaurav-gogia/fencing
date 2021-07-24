package main

import "fmt"

func concat_delim() {
	var joinedStrings []string

	delim := "Hi"
	testCase := []string{"Saurabh", "Hi", "How", "are", "you", "Hi", "I'm", "fine"}
	// [“Saurabh”,"Hi How are you", "Hi I'm fine"]

	for outerIndex := 0; outerIndex < len(testCase); outerIndex++ {

		if testCase[outerIndex] == delim {
			var tmp string

			tmp += testCase[outerIndex] + " "
			for innerIndex := outerIndex + 1; innerIndex < len(testCase); innerIndex++ {
				if testCase[innerIndex] == delim {
					outerIndex = innerIndex - 1
					break
				}

				tmp += testCase[innerIndex] + " "

				if innerIndex == len(testCase)-1 {
					outerIndex = innerIndex
					break
				}
			}

			joinedStrings = append(joinedStrings, tmp)
		} else {
			joinedStrings = append(joinedStrings, testCase[outerIndex])
		}
	}

	fmt.Println(joinedStrings)
}
