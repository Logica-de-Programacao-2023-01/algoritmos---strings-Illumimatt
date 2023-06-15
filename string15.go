package main

import (
	"fmt"
	"strings"
)

func replaceVowels(input string) string {
	vowels := "aeiouAEIOU"
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(vowels, r) {
			return '*'
		}
		return r
	}, input)
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	modifiedString := replaceVowels(input)
	fmt.Println("String com vogais substituídas:", modifiedString)
}
