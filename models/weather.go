package models

import "github.com/shopspring/decimal"

type WeathersForecast struct {
	Count int               `json:"count"`
	Data  []WeatherForecast `json:"data"`
}

type Weather struct {
	Code        int    `json:"code"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

type WeatherForecast struct {
	AppTemp      decimal.Decimal `json:"app_temp"`
	Aqi          decimal.Decimal `json:"aqi"`
	CityName     string          `json:"city_name"`
	Clouds       decimal.Decimal `json:"clouds"`
	CountryCode  string          `json:"country_code"`
	Datetime     string          `json:"datetime"`
	Dewpt        decimal.Decimal `json:"dewpt"`
	Dhi          decimal.Decimal `json:"dhi"`
	Dni          decimal.Decimal `json:"dni"`
	ElevAngle    decimal.Decimal `json:"elev_angle"`
	Ghi          decimal.Decimal `json:"ghi"`
	Gust         interface{}     `json:"gust"`
	HAngle       decimal.Decimal `json:"h_angle"`
	Lat          decimal.Decimal `json:"lat"`
	Lon          decimal.Decimal `json:"lon"`
	ObTime       string          `json:"ob_time"`
	Pod          string          `json:"pod"`
	Precip       decimal.Decimal `json:"precip"`
	Pres         decimal.Decimal `json:"pres"`
	Rh           decimal.Decimal `json:"rh"`
	Slp          decimal.Decimal `json:"slp"`
	Snow         decimal.Decimal `json:"snow"`
	SolarRad     decimal.Decimal `json:"solar_rad"`
	Sources      []string        `json:"sources"`
	StateCode    string          `json:"state_code"`
	Station      string          `json:"station"`
	Sunrise      string          `json:"sunrise"`
	Sunset       string          `json:"sunset"`
	Temp         decimal.Decimal `json:"temp"`
	Timezone     string          `json:"timezone"`
	Ts           decimal.Decimal `json:"ts"`
	Uv           decimal.Decimal `json:"uv"`
	Vis          decimal.Decimal `json:"vis"`
	Weather      Weather         `json:"weather"`
	WindCdir     string          `json:"wind_cdir"`
	WindCdirFull string          `json:"wind_cdir_full"`
	WindDir      decimal.Decimal `json:"wind_dir"`
	WindSpd      decimal.Decimal `json:"wind_spd"`
}
