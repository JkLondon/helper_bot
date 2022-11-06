package httpClient

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"weatherEveryDay/models"

	"github.com/shopspring/decimal"
)

func GetWeather(apiKey string, lon decimal.Decimal, lat decimal.Decimal) (weatherForecast models.WeathersForecast, err error) {
	url := fmt.Sprintf(CurrentWeather, lon.String(), lat.String())
	err = MakeRequest(models.RequestParams{
		Url:    url,
		Host:   WeatherHost,
		ApiKey: apiKey,
		Dest:   &weatherForecast,
	})
	if err != nil {
		return weatherForecast, err
	}
	return weatherForecast, nil
}

func GetCity(apiKey string, cityQuery string) (cities []models.CityInfo, err error) {
	url := fmt.Sprintf(SearchCity, cityQuery)
	err = MakeRequest(models.RequestParams{
		Url:    url,
		Host:   CityHost,
		ApiKey: apiKey,
		Dest:   &cities,
	})
	if err != nil {
		return cities, err
	}
	return cities, nil
}

func MakeRequest(params models.RequestParams) (err error) {
	req, err := http.NewRequest("GET", params.Url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Key", params.ApiKey)
	req.Header.Add("X-RapidAPI-Host", params.Host)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	println(string(body))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &params.Dest)
	if err != nil {
		return err
	}
	return nil
}
