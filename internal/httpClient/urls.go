package httpClient

const (
	CurrentWeather = "https://weatherbit-v1-mashape.p.rapidapi.com/current?lon=%s&lat=%s&lang=ru"
	WeatherHost    = "weatherbit-v1-mashape.p.rapidapi.com"

	SearchCity = "https://spott.p.rapidapi.com/places/autocomplete?type=CITY&q=%s&country=RU&skip=0"
	CityHost   = "spott.p.rapidapi.com"

	JiraHistoryRequest = "https://cbgamma.atlassian.net/rest/api/2/search?startAt=%d"
	JiraIssueChangelog = "https://cbgamma.atlassian.net/rest/api/2/issue/%s/changelog"
)
