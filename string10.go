package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(str1 string, str2 string) bool {
	str1 = strings.ReplaceAll(strings.ToLower(str1), " ", "")
	str2 = strings.ReplaceAll(strings.ToLower(str2), " ", "")
	str1Runes := []rune(str1)
	str2Runes := []rune(str2)
	sort.Slice(str1Runes, func(i, j int) bool { return str1Runes[i] < str1Runes[j] })
	sort.Slice(str2Runes, func(i, j int) bool { return str2Runes[i] < str2Runes[j] })
	return string(str1Runes) == string(str2Runes)
}

func main() {
	var str1, str2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	if isAnagram(str1, str2) {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}
