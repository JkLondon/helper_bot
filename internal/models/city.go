package models

import "github.com/shopspring/decimal"

type CityInfo struct {
	Id         string          `json:"id"`
	GeonameId  decimal.Decimal `json:"geonameId"`
	Type       string          `json:"type"`
	Name       string          `json:"name"`
	Population decimal.Decimal `json:"population"`
	Elevation  decimal.Decimal `json:"elevation"`
	TimezoneId string          `json:"timezoneId"`
	Country    struct {
		Id        string          `json:"id"`
		GeonameId decimal.Decimal `json:"geonameId"`
		Name      string          `json:"name"`
	} `json:"country"`
	AdminDivision1 struct {
		Id        string          `json:"id"`
		GeonameId decimal.Decimal `json:"geonameId"`
		Name      string          `json:"name"`
	} `json:"adminDivision1"`
	Score       decimal.Decimal `json:"score"`
	Coordinates struct {
		Latitude  decimal.Decimal `json:"latitude"`
		Longitude decimal.Decimal `json:"longitude"`
	} `json:"coordinates"`
}
