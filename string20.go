package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isCamelCase(input string) bool {
	// Verificar se a primeira letra é minúscula
	if len(input) > 0 && unicode.IsUpper(rune(input[0])) {
		return false
	}

	// Verificar se há letras minúsculas após a primeira letra de cada palavra
	for i := 1; i < len(input); i++ {
		if unicode.IsUpper(rune(input[i])) && unicode.IsLower(rune(input[i-1])) {
			return false
		}
	}

	return true
}

func countWord(input string) int {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return unicode.IsUpper(r) || unicode.IsSpace(r)
	})

	return len(words)
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isCamelCase(input) {
		wordCount := countWord(input)
		fmt.Println("A string está em camelCase.")
		fmt.Println("Quantidade de palavras:", wordCount)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}
