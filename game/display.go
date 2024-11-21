package game

import "strings"

func DisplayWord(word string, guessedLetters map[rune]bool) string {
	var displayed strings.Builder
	for _, letter := range word {
		if guessedLetters[letter] {
			displayed.WriteRune(letter)
		} else {
			displayed.WriteRune('_')
		}
		displayed.WriteRune(' ')
	}
	return displayed.String()
}
