package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}

func main() {
	var input string
	fmt.Print("Digite uma frase: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	numWords := countWords(input)
	fmt.Printf("A frase contém %d palavras.\n", numWords)
}
