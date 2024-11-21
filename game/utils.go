package game

func AllLettersGuessed(word string, guessedLetters map[rune]bool) bool {
	for _, letter := range word {
		if !guessedLetters[letter] {
			return false
		}
	}
	return true
}
