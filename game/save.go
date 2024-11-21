package game

import (
	"encoding/json"
	"os"
)

func SaveGame(state GameState, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(state)
}
