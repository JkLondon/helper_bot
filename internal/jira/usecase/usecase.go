package usecase

import (
	"fmt"
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
		task := models.Task{
			Name:        issue.Fields.Summary,
			Assignee:    assignee,
			Assignees:   nil,
			Description: desc,
			Peredogovor: 0,
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
	for i, task := range jiraData.Tasks {
		result += fmt.Sprintf("Задача №%d: %s. %s\n", i, task.Name, task.Assignee)
	}
	return result, nil
}
