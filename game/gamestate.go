package game

type GameState struct {
	Word           string        `json:"word"`
	GuessedLetters map[rune]bool `json:"guessed_letters"`
	RemainingTries int           `json:"remaining_tries"`
}
