package main

import (
	"fmt"
	strUtil "strings"
)

func main() {
	word := "Превет"
	fmt.Println("source:", word)
	
	// solition 1
	fmt.Println("target:", strUtil.Replace(word, "е", "и", 1))

	// solution 2 (with rune slice)
	wordSlice := make([]rune, 0)
	for _, runeValue := range word {
		wordSlice = append(wordSlice, runeValue)
    }
	wordSlice[2] = 'и'
	fmt.Println("target:", string(wordSlice))
}
