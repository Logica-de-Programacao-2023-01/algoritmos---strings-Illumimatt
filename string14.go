package main

import (
	"fmt"
	"strconv"
)

func isNumericDecrescense(input string) bool {
	// Converter a string em um slice de runes
	runes := []rune(input)

	// Verificar se a sequência é numérica decrescente
	for i := 1; i < len(runes); i++ {
		curr, _ := strconv.Atoi(string(runes[i]))
		prev, _ := strconv.Atoi(string(runes[i-1]))

		if curr != prev-1 {
			return false
		}
	}

	return true
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isNumericDecrescense(input) {
		fmt.Println("É uma sequência numérica decrescente.")
	} else {
		fmt.Println("Não é uma sequência numérica decrescente.")
	}
}
