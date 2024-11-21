package game

import (
	"fmt"
	"os"
	"strings"
)

func PlayHangman(word string, guessedLetters map[rune]bool, remainingTries int) {
	var guess string

	for remainingTries > 0 {
		DisplayHangman(10 - remainingTries)
		fmt.Println("\nMot à deviner:", DisplayWord(word, guessedLetters))
		fmt.Printf("Il vous reste %d essais. Tapez une lettre, un mot complet ou STOP pour sauvegarder : ", remainingTries)

		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Erreur de lecture. Réessayez.")
			continue
		}

		if strings.ToUpper(guess) == "STOP" {
			state := GameState{
				Word:           word,
				GuessedLetters: guessedLetters,
				RemainingTries: remainingTries,
			}
			err := SaveGame(state, "save.txt")
			if err != nil {
				fmt.Println("Erreur lors de la sauvegarde de la partie :", err)
			} else {
				fmt.Println("Partie sauvegardée dans save.txt. À bientôt !")
			}
			return
		}

		if len(guess) > 1 {
			if strings.ToLower(guess) == strings.ToLower(word) {
				fmt.Println("\nFélicitations ! Vous avez deviné le mot :", word)
				deleteSaveFile("save.txt")
				return
			} else {
				remainingTries -= 2
				fmt.Println("Mauvaise réponse. Vous perdez 2 vies.")
				continue
			}
		}

		if len(guess) != 1 {
			fmt.Println("Entrez une seule lettre ou un mot complet.")
			continue
		}

		letter := rune(guess[0])

		if guessedLetters[letter] {
			fmt.Println("Vous avez déjà deviné cette lettre.")
			continue
		}

		guessedLetters[letter] = true

		if !strings.ContainsRune(word, letter) {
			remainingTries--
			fmt.Println("Mauvaise réponse.")
		} else {
			fmt.Println("Bonne réponse !")
		}

		if AllLettersGuessed(word, guessedLetters) {
			fmt.Println("\nFélicitations ! Vous avez deviné le mot :", word)
			deleteSaveFile("save.txt")
			return
		}
	}

	DisplayHangman(10)
	fmt.Println("\nVous avez perdu ! Le mot était :", word)
	deleteSaveFile("save.txt")
}

func deleteSaveFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du fichier de sauvegarde :", err)
	}
}
