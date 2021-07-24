package main

import "fmt"

type testCase struct {
	caseValue    []string
	caseExpected bool
}

func match_count() {
	name := "mango"

	testCase1 := testCase{
		caseValue:    []string{"man", "go", "ang"},
		caseExpected: true,
	}
	testCase2 := testCase{
		caseValue:    []string{"man", "go", "ang", "a", "n", "ng"},
		caseExpected: false,
	}
	testCase3 := testCase{
		caseValue:    []string{"man", "go", "ang", "a", "n", "gn"},
		caseExpected: true,
	}
	testCase4 := testCase{
		caseValue:    []string{"mn", "goo", "ang", "a", "n"},
		caseExpected: true,
	}

	testCases := map[int]testCase{
		1: testCase1,
		2: testCase2,
		3: testCase3,
		4: testCase4,
	}

	for _, testCaseObject := range testCases {
		if testCaseObject.caseExpected == matcher(testCaseObject.caseValue, name) {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}
	}
}

func matcher(testCase []string, name string) bool {
	counter := 0

	for _, str := range testCase {
		if matchConstraints(str, name) {
			counter++
		}
	}

	return (counter >= 2) && (counter <= 5)
}

func matchConstraints(testString, nameString string) bool {
	counter := 0
	j := -1

	for i := 0; i < len(testString); i++ {

		for j < len(nameString) {
			j++

			if j == len(nameString) {
				break
			}

			if testString[i] == nameString[j] {
				counter++
				break
			} else {
				i = 0
			}

		}

		if j == len(nameString) {
			break
		}
	}

	return counter == len(testString)
}
