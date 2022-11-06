package templates

import (
	"fmt"
	"weatherEveryDay/models"
)

const (
	TgMsg = "Погода в г. %s: %s\nТемпература: %s, Ощущается как %s.\nВетер %s. Движется со скоростью %s м/c\nВосход в %s, закат в %s.\n"
)

func MakeTGWeatherMessage(weather models.WeatherForecast) string {
	return fmt.Sprintf(
		TgMsg,
		weather.CityName,
		weather.Weather.Description,
		weather.Temp.String(),
		weather.AppTemp.String(),
		weather.WindCdirFull,
		weather.WindSpd.String(),
		weather.Sunrise,
		weather.Sunset,
	)
}
