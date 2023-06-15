package main

import (
	"fmt"
	"strings"
)

func removeVowels(input string) string {
	vowels := "aeiouAEIOU"
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(vowels, r) {
			return -1 // Remove a vogal
		}
		return r
	}, input)
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := removeVowels(input)
	fmt.Println("String sem vogais:", result)
}
