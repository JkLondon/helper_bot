package jira

import (
	"weatherEveryDay/internal/models"
)

type UseCase interface {
	ParseRawDataToFocusReport(params models.JiraRawData) (result models.JiraFocusData, err error)
	MakeFocusReport(params models.JiraRawData) (result string, err error)
	ParseRawDataToDailyReport(params models.JiraRawData) (result models.JiraDailyData, err error)
	MakeDailyReport(params models.JiraRawData) (result string, err error)
}
