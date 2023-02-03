package models

type RequestRapidAPIParams struct {
	Url    string
	Host   string
	ApiKey string
	Dest   interface{}
}

type RequestJiraParams struct {
	Url   string
	Login string
	Token string
	Dest  interface{}
}
