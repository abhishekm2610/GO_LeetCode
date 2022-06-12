package main

import (
	"fmt"
	"strings"
)

//Original Solution - Took too much memory
func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

//Optimizing Solution - using iterative approach
func defangIPaddrOp(address string) string {
	deFangedString := ""
	for i := 0; i < len(address); i++ {
		if address[i] != '.' {
			deFangedString += string(address[i])
		} else {
			deFangedString += "[.]"
		}
	}
	return deFangedString
}

func main() {
	input := "1.1.1.1"

	fmt.Println(defangIPaddrOp((input)))
}
