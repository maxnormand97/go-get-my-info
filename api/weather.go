package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"maxnormand/get-my-info/models"
)

// init function to pre-load the .env file
func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetWeather() (string, error) {
	fmt.Println("Fetching the Weather...")
	// Load the API key from the environment variable
	apiKey := os.Getenv("WEATHERSTACK_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("API key not found in environment variables")
	}

	// Construct the API URL using the API key
	apiURL := fmt.Sprintf("http://api.weatherstack.com/current?access_key=%s&query=37.8267,-122.4233", apiKey)

	res, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("weather API down")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var weatherResponse models.WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return "", err
	}

	// Access data and format the string response
	response := fmt.Sprintf(
		"Location: %s, %s, %s\nTemperature: %d°C\nWeather: %s\nWind: %d km/h %s\n",
		weatherResponse.Location.Name,
		weatherResponse.Location.Region,
		weatherResponse.Location.Country,
		weatherResponse.Current.Temperature,
		weatherResponse.Current.WeatherDescriptions[0],
		weatherResponse.Current.WindSpeed,
		weatherResponse.Current.WindDir,
	)

	return response, nil
}
