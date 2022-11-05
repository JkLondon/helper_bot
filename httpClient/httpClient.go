package httpClient

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"weatherEveryDay/models"
)

func GetWeather(apiKey string) (weatherForecast models.WeathersForecast, err error) {
	url := fmt.Sprintf(CurrentWeather, LonDolgoprudnyy, LatDolgoprudnyy)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return weatherForecast, err
	}

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", WeatherHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return weatherForecast, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return weatherForecast, err
	}
	err = json.Unmarshal(body, &weatherForecast)
	if err != nil {
		return weatherForecast, err
	}
	return weatherForecast, nil
}
