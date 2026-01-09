package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	if len(word) <= 1 {
		return word
	}

	firstChar := word[:1]
	rest := word[1:]

	runes := []rune(rest)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return firstChar + string(runes)
}

func encryptPhrase(phrase string) string {
	words := strings.Split(phrase, " ")

	for i, word := range words {
		words[i] = encryptWord(word)
	}

	return strings.Join(words, " ")
}

func main() {
	phrases := []string{
		"Pepe Schnele is a legend",
		"Hello world",
		"Go is cool",
	}

	for _, phrase := range phrases {
		fmt.Println("Исходная фраза:", phrase)
		fmt.Println("Зашифрованная:", encryptPhrase(phrase))
		fmt.Println()
	}
}
