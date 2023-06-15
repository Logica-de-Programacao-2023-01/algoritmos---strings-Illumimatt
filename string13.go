package main

import (
	"fmt"
	"strconv"
)

func isNumericSequence(input string) bool {
	// Converter a string em um slice de runes
	runes := []rune(input)

	// Verificar se a sequência é numérica crescente
	for i := 1; i < len(runes); i++ {
		curr, _ := strconv.Atoi(string(runes[i]))
		prev, _ := strconv.Atoi(string(runes[i-1]))

		if curr != prev+1 {
			return false
		}
	}

	return true
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isNumericSequence(input) {
		fmt.Println("É uma sequência numérica crescente.")
	} else {
		fmt.Println("Não é uma sequência numérica crescente.")
	}
}
