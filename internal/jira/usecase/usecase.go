package usecase

import (
	"fmt"
	"time"
	"weatherEveryDay/internal/jira"
	"weatherEveryDay/internal/models"
)

type JiraUC struct {
}

func NewJiraUC() jira.UseCase {
	return &JiraUC{}
}

func (j *JiraUC) ParseRawData(params models.JiraRawData) (result models.JiraData, err error) {
	result.TotalIssues = params.Total
	result.TotalMonth = 0
	result.TotalWeek = 0
	result.Tasks = make([]models.Task, 0)
	for _, issue := range params.Issues {
		desc := ""
		if issue.Fields.Description != nil {
			desc = *issue.Fields.Description
		}
		assignee := ""
		if issue.Fields.Assignee != nil {
			assignee = issue.Fields.Assignee.DisplayName
		}

		shortForm := "2023-02-02"
		task := models.Task{
			Name:        issue.Fields.Summary,
			Assignee:    assignee,
			Assignees:   nil,
			Description: desc,
			Peredogovor: 0,
			Status:      issue.Fields.Status.Name,
		}
		if issue.Fields.Duedate != nil {
			task.DueTo, err = time.Parse(shortForm, *issue.Fields.Duedate)
			if err != nil {
				return result, err
			}
		}
		if task.Status == "Готово" && issue.Fields.Duedate != nil {
			if task.DueTo.Add(time.Hour * 24 * 30).After(time.Now()) {
				result.TotalMonth += 1
			}
			if task.DueTo.Add(time.Hour * 24 * 7).After(time.Now()) {
				result.TotalMonth += 1
			}
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result, nil
}

func (j *JiraUC) MakeReport(params models.JiraRawData) (result string, err error) {
	jiraData, err := j.ParseRawData(params)
	if err != nil {
		return "", err
	}
	result += fmt.Sprintf("Всего задач %d\n", jiraData.TotalIssues)
	result += fmt.Sprintf("За месяц было выполнено %d задач, за неделю - %d\n", jiraData.TotalMonth, jiraData.TotalWeek)
	return result, nil
}
