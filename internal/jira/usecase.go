package jira

import (
	"weatherEveryDay/internal/models"
)

type UseCase interface {
	ParseRawData(params models.JiraRawData) (result models.JiraData, err error)
	MakeReport(params models.JiraRawData) (result string, err error)
}
