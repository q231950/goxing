// remove_punctuation.go
package main

import (
	"unicode"
	"fmt"
	"strings"
)
type RuneForRuneFunc func(rune) rune

func main() {
	var removePunctuation RuneForRuneFunc

	phrases := []string{"Day; duns. and night", "All day long"}
	removePunctuation = func(char rune) rune {
		if unicode.Is(unicode.Terminal_Punctuation, char) {
			return -1
		}
		return char
	}

	processPrases(phrases, removePunctuation)
}

func processPrases(phrases []string, function RuneForRuneFunc) {
	for _, phrase := range phrases {
		fmt.Println(strings.Map(function, phrase))
	}
}
