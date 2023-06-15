package main

import (
	"fmt"
	"strings"
)

func replaceLetter(input string, oldLetter rune, newLetter rune) string {
	return strings.ReplaceAll(input, string(oldLetter), string(newLetter))
}

func main() {
	var input string
	var oldLetter, newLetter rune

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Print("Digite a letra que deseja substituir: ")
	fmt.Scanf("%c", &oldLetter)

	fmt.Print("Digite a letra de substituição: ")
	fmt.Scanf("%c", &newLetter)

	modifiedString := replaceLetter(input, oldLetter, newLetter)
	fmt.Println("String modificada:", modifiedString)
}
