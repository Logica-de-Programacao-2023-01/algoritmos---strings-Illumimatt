package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba uma string e substitua todas as ocorrências desse caractere na string
// por outro caractere especificado pelo usuário.
// Informe ao usuário o resultado.
func main() {
	var x, y string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma string: ")
	texto, _ := reader.ReadString('\n')
	fmt.Scan(&x)
	fmt.Scan(&y)
	textofinal := strings.ReplaceAll(texto, x, y)
	fmt.Print(textofinal)
}
