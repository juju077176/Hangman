package game

import (
	"encoding/json"
	"os"
)

func LoadGame(filePath string) (GameState, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return GameState{}, err
	}
	defer file.Close()

	var state GameState
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&state)
	if err != nil {
		return GameState{}, err
	}

	return state, nil
}
