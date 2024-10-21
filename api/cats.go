package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"maxnormand/get-my-info/models"
)

func GetCatFacts() (string, error) {
	res, err := http.Get("https://cat-fact.herokuapp.com/facts")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("cat API down")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var todaysCatFact []models.CatFact
	err = json.Unmarshal([]byte(body), &todaysCatFact)
	if err != nil {
		return "", err
	}

	// TODO: here need to handle a better way to format the req from the chess masters

	return todaysCatFact[0].Text, nil
}
