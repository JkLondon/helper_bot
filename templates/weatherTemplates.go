package templates

import (
	"fmt"
	"weatherEveryDay/models"
)

const (
	TgMsg = "Погода в г. %s: %s\nТемпература: %d, Ощущается как %d.\nВетер %s. Движется со скоростью %f м/c\nВосход в %s, закат в %s.\n"
)

func MakeTGWeatherMessage(weather models.WeatherForecast) string {
	return fmt.Sprintf(
		TgMsg,
		weather.CityName,
		weather.Weather.Description,
		weather.Temp,
		int(weather.AppTemp),
		weather.WindCdirFull,
		weather.WindSpd,
		weather.Sunrise,
		weather.Sunset,
	)
}
