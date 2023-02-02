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
	Pass  string
	Dest  interface{}
}
