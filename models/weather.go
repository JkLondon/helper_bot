package models

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
	AppTemp      float64     `json:"app_temp"`
	Aqi          int         `json:"aqi"`
	CityName     string      `json:"city_name"`
	Clouds       int         `json:"clouds"`
	CountryCode  string      `json:"country_code"`
	Datetime     string      `json:"datetime"`
	Dewpt        float64     `json:"dewpt"`
	Dhi          int         `json:"dhi"`
	Dni          int         `json:"dni"`
	ElevAngle    float64     `json:"elev_angle"`
	Ghi          int         `json:"ghi"`
	Gust         interface{} `json:"gust"`
	HAngle       int         `json:"h_angle"`
	Lat          float64     `json:"lat"`
	Lon          float64     `json:"lon"`
	ObTime       string      `json:"ob_time"`
	Pod          string      `json:"pod"`
	Precip       int         `json:"precip"`
	Pres         float64     `json:"pres"`
	Rh           int         `json:"rh"`
	Slp          float64     `json:"slp"`
	Snow         int         `json:"snow"`
	SolarRad     int         `json:"solar_rad"`
	Sources      []string    `json:"sources"`
	StateCode    string      `json:"state_code"`
	Station      string      `json:"station"`
	Sunrise      string      `json:"sunrise"`
	Sunset       string      `json:"sunset"`
	Temp         int         `json:"temp"`
	Timezone     string      `json:"timezone"`
	Ts           int         `json:"ts"`
	Uv           int         `json:"uv"`
	Vis          int         `json:"vis"`
	Weather      Weather     `json:"weather"`
	WindCdir     string      `json:"wind_cdir"`
	WindCdirFull string      `json:"wind_cdir_full"`
	WindDir      int         `json:"wind_dir"`
	WindSpd      float64     `json:"wind_spd"`
}
