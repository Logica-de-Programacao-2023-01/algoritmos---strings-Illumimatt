package main

import (
	"fmt"
	"strings"
)

// Escreva um programa que
// receba uma string e converta todas as letras minúsculas para maiúsculas.
// Informe ao usuário o resultado.
func main() {
	var texto string
	fmt.Scan(&texto)
	fmt.Print(strings.ToUpper(texto))
}
