package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba uma string e remova todos os espaços em branco.
// Informe ao usuário o resultado.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma string: ")
	str, _ := reader.ReadString('\n')
	strSemEspacos := strings.ReplaceAll(str, " ", "")
	fmt.Println("String original:", str)
	fmt.Println("String sem espaços:", strSemEspacos)
}
