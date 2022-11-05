package main

import (
	"fmt"
	"log"
	"os"
	"weatherEveryDay/httpClient"
	"weatherEveryDay/templates"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	var apiKey = os.Getenv("WEATHER_API_KEY")
	weatherForecast, err := httpClient.GetWeather(apiKey)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	weather := weatherForecast.Data[0]
	message := templates.MakeTGWeatherMessage(weather)
	fmt.Println(message)
}
