package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"maxnormand/get-my-info/models"
)

func GetChessPlayers() ([]string, error) {
	res, err := http.Get("https://api.chess.com/pub/titled/GM")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("chess API down")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var topChessPlayers models.ChessPlayers
	err = json.Unmarshal(body, &topChessPlayers)
	if err != nil {
		return nil, err
	}

	// TODO: here need to handle a better way to format the req from the chess masters

	return topChessPlayers.Players, nil
}
