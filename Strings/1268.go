package main

import (
	"fmt"
	"sort"
)

// Original Solution
// Runtime: 54 ms
// Memory Usage: 10.5 MB
func suggestedProducts(products []string, searchWord string) [][]string {
	outputArray := [][]string{}
	var searchString = ""
	for _, char := range searchWord {
		searchString += string(char)
		listOfMatches := []string{}
		for _, value := range products {
			if len(value) >= len(searchString) {
				if searchString == value[:len(searchString)] {
					listOfMatches = append(listOfMatches, string(value))
				}
			} else if value == searchString {
				listOfMatches = append(listOfMatches, string(value))
			}

		}
		sort.Strings(listOfMatches)
		if len(listOfMatches) < 3 {
			outputArray = append(outputArray, listOfMatches)
		} else {
			outputArray = append(outputArray, listOfMatches[:3])
		}
	}
	return outputArray
}

func main() {
	input1 := []string{"bags", "baggage", "banner", "box", "cloths"}
	input2 := "bags"

	fmt.Println(suggestedProducts(input1, input2))
}
