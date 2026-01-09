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
	runes := []rune(word[1:])

	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
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
