package main

import (
	"fmt"
	"regexp"
)

func containsNumber(input string) bool {
	match, _ := regexp.MatchString(`\d`, input)
	return match
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if containsNumber(input) {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém números.")
	}
}
