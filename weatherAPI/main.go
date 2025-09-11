package main 

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is not set")
	}

	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=Tehran&appid=%s", apiKey)

	response, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal("Error fetching weather data")
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)	

	var weatherData map[string]interface{}
	err = json.Unmarshal(responseBody, &weatherData)
	if err != nil {
		log.Fatal("Error unmarshalling weather data")
	}

	fmt.Printf("Weather in Tehran: %+v\n", weatherData)
}