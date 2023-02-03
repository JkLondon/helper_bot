package httpClient

import (
	"encoding/json"
	"fmt"
	"weatherEveryDay/internal/models"

	"io/ioutil"
	"net/http"

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

func FetchRawJiraHistory(
	login string,
	token string,
) (result models.JiraRawData, err error) {
	startAt := 0
	err = FetchJiraData(models.RequestJiraParams{
		Url:   fmt.Sprintf(JiraHistoryRequest, startAt),
		Login: login,
		Token: token,
		Dest:  &result,
	})
	if err != nil {
		return result, err
	}
	startAt += result.MaxResults
	for startAt < result.Total {
		tmpRes := models.JiraRawData{}
		err = FetchJiraData(models.RequestJiraParams{
			Url:   fmt.Sprintf(JiraHistoryRequest, startAt),
			Login: login,
			Token: token,
			Dest:  &tmpRes,
		})
		if err != nil {
			return result, err
		}
		startAt += result.MaxResults
		result.Issues = append(result.Issues, tmpRes.Issues...)
	}
	return result, nil
}

func FetchIssueChangelog(
	login string,
	token string,
	issueID string,
) (result models.JiraRawChangelog, err error) {
	err = FetchJiraData(models.RequestJiraParams{
		Url:   fmt.Sprintf(JiraIssueChangelog, issueID),
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

func FetchJiraData(params models.RequestJiraParams) (err error) {
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
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &params.Dest)
	if err != nil {
		return err
	}
	return nil
}
