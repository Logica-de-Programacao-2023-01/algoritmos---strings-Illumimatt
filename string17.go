package main

import (
	"fmt"
)

func getUniqueLetters(input string) string {
	// Mapa para contar a ocorrência de cada letra
	letterCount := make(map[rune]int)

	// Contar a ocorrência de cada letra na string
	for _, char := range input {
		letterCount[char]++
	}

	// Construir uma nova string apenas com as letras únicas
	uniqueLetters := ""
	for _, char := range input {
		if letterCount[char] == 1 {
			uniqueLetters += string(char)
		}
	}

	return uniqueLetters
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	unique := getUniqueLetters(input)
	fmt.Println("Letras únicas na string:", unique)
}
