package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
	// Remover espaços em branco e converter para letras minúsculas
	input = strings.ReplaceAll(strings.ToLower(input), " ", "")

	// Verificar se a string é um palíndromo
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("É um palíndromo.")
	} else {
		fmt.Println("Não é um palíndromo.")
	}
}
