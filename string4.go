package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma string: ")
	texto, _ := reader.ReadString('\n')

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma string: ")
	texto2, _ := reader2.ReadString('\n')
	resultado := strings.Compare(texto, texto2)
	if resultado == 0 {
		fmt.Println("As strings são iguais.")
	} else {
		fmt.Println("As strings diferem.")
	}
}
