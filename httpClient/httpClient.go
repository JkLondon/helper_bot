package httpClient

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"weatherEveryDay/models"

	"github.com/shopspring/decimal"
)

func GetWeather(
	apiKey string,
	lon decimal.Decimal,
	lat decimal.Decimal,
) (weatherForecast models.WeathersForecast, err error) {
	url := fmt.Sprintf(CurrentWeather, lon.String(), lat.String())
	err = MakeRapidAPIRequest(models.RequestRapidAPIParams{
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

func GetCity(
	apiKey string,
	cityQuery string,
) (cities []models.CityInfo, err error) {
	url := fmt.Sprintf(SearchCity, cityQuery)
	err = MakeRapidAPIRequest(models.RequestRapidAPIParams{
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

func GetJiraReport(
	login string,
	token string,
) (result models.JiraReport, err error) {
	err = MakeJiraRequest(models.RequestJiraParams{
		Url:   JiraReq,
		Login: login,
		Token: token,
		Dest:  &result,
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

func MakeRapidAPIRequest(params models.RequestRapidAPIParams) (err error) {
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

func MakeJiraRequest(params models.RequestJiraParams) (err error) {
	req, err := http.NewRequest("GET", params.Url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(params.Login, params.Token)

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
