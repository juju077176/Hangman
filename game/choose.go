package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func ChooseRandomWord(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(words) == 0 {
		return "", fmt.Errorf("le fichier est vide")
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	return words[randomIndex], nil
}
