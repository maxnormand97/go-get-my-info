package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"maxnormand/get-my-info/models"
)

func GetChessPuzzle() (string, error) {
	fmt.Println("Fetching the Daily Chess puzzle...")
	res, err := http.Get("https://api.chess.com/pub/puzzle")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("chess API down")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var dailyPuzzleRes models.DailyPuzzleResponse
	err = json.Unmarshal(body, &dailyPuzzleRes)
	if err != nil {
		return "", err
	}

	return dailyPuzzleRes.URL, nil
}
