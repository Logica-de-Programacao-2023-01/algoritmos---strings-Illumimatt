package main

import (
	"fmt"
	"strings"
)

func isSubstring(mainString, substring string) bool {
	return strings.Contains(mainString, substring)
}

func main() {
	var mainString, substring string

	fmt.Print("Digite a string principal: ")
	fmt.Scanln(&mainString)

	fmt.Print("Digite a substring: ")
	fmt.Scanln(&substring)

	if isSubstring(mainString, substring) {
		fmt.Println("A segunda string é uma substring da primeira.")
	} else {
		fmt.Println("A segunda string não é uma substring da primeira.")
	}
}
