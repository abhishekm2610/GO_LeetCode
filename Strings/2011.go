package main

import "fmt"

//Original Solution
func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, value := range operations {
		if value == "--X" || value == "X--" {
			x -= 1
		} else {
			x += 1
		}
	}
	return x
}

//Trying out Optimal One
func finalValueAfterOperationsOp(operations []string) int {
	x := 0
	for _, value := range operations {
		if value[1] == '-' {
			x--
		} else {
			x++
		}
	}
	return x
}

func main() {
	input := []string{"X++", "++X", "--X", "X--"}

	fmt.Println(finalValueAfterOperationsOp((input)))
}
