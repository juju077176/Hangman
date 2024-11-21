package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/juju077176/Hangman/game"
)

func main() {
	startWithSave := flag.String("startWith", "", "Reprendre une partie sauvegardée avec le fichier spécifié.")
	flag.Parse()

	var gameState game.GameState
	var err error

	if *startWithSave != "" {
		fmt.Println("Reprise de la partie à partir du fichier :", *startWithSave)
		gameState, err = game.LoadGame(*startWithSave)
		if err != nil {
			fmt.Println("Erreur lors du chargement de la partie :", err)
			fmt.Println("Démarrage d'une nouvelle partie à la place.")
			startNewGame()
			return
		}
	} else if _, err := os.Stat("save.txt"); err == nil {
		fmt.Println("Reprise automatique de la partie à partir de save.txt...")
		gameState, err = game.LoadGame("save.txt")
		if err != nil {
			fmt.Println("Erreur lors du chargement de la partie :", err)
			fmt.Println("Démarrage d'une nouvelle partie à la place.")
			startNewGame()
			return
		}
	} else {
		startNewGame()
		return
	}

	game.PlayHangman(gameState.Word, gameState.GuessedLetters, gameState.RemainingTries)
}

func startNewGame() {
	word, err := game.ChooseRandomWord("words.txt")
	if err != nil {
		fmt.Println("Erreur lors du chargement de la liste des mots :", err)
		return
	}

	guessedLetters := make(map[rune]bool)

	firstLetter := rune(word[0])
	guessedLetters[firstLetter] = true

	remainingTries := 10

	game.PlayHangman(word, guessedLetters, remainingTries)
}
