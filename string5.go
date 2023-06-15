package main

import (
	"fmt"
	"strconv"
)

//Escreva um programa que receba uma string
//e verifique se ela é um número válido em ponto flutuante.
//Informe ao usuário se é ou não.

func isFloatNumber(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}

func main() {
	var input string
	fmt.Print("Digite um número em ponto flutuante: ")
	fmt.Scanln(&input)

	if isFloatNumber(input) {
		fmt.Println("É um número válido em ponto flutuante.")
	} else {
		fmt.Println("Não é um número válido em ponto flutuante.")
	}
}
