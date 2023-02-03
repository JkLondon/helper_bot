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
	result.MemberMapTasksDoneWeek = make(map[string]int)
	result.MemberMapTasksDoneMonth = make(map[string]int)
	result.MemberMapPeredogovorsWeek = make(map[string]int)
	result.MemberMapPeredogovorsMonth = make(map[string]int)
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
			Assignees:   nil,
			Description: desc,
			Peredogovor: 0,
			Status:      issue.Fields.Status.Name,
		}

		if assignee != "" {
			task.Assignees = append(task.Assignees, assignee)
		}

		for _, mem := range issue.Fields.Workers {
			task.Assignees = append(task.Assignees, mem.DisplayName)
		}

		if issue.Fields.Duedate != nil {
			shortForm := "2006-01-02"
			dueDateStr := *issue.Fields.Duedate
			task.DueTo, err = time.Parse(shortForm, dueDateStr)
			if err != nil {
				return result, err
			}
		}
		if task.Status == "Готово" && issue.Fields.Duedate != nil {
			if task.DueTo.Add(time.Hour * 24 * 30).After(time.Now()) {
				result.TotalMonth += 1
				for _, name := range task.Assignees {
					result.MemberMapTasksDoneMonth[name] += 1
				}
			}
			if task.DueTo.Add(time.Hour * 24 * 7).After(time.Now()) {
				result.TotalWeek += 1
				for _, name := range task.Assignees {
					result.MemberMapTasksDoneWeek[name] += 1
				}
			}
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result, nil
}

func (j *JiraUC) MakeReport(params models.JiraRawData) (result string, err error) {
	jiraData, err := j.ParseRawData(params)
	if err != nil {
		return err.Error(), err
	}
	result += fmt.Sprintf("Всего задач %d\n", jiraData.TotalIssues)
	result += fmt.Sprintf("За месяц было выполнено %d задач, за неделю — %d\n", jiraData.TotalMonth, jiraData.TotalWeek)
	result += fmt.Sprintf("За эту неделю %d cотрудников выполнили определенное количество задач, а именно:\n", len(jiraData.MemberMapTasksDoneWeek))
	for key, value := range jiraData.MemberMapTasksDoneWeek {
		result += fmt.Sprintf("%s выполнил(а) %d задач\n", key, value)
	}
	result += fmt.Sprintf("За этот месяц %d cотрудников выполнили определенное количество задач, а именно:\n", len(jiraData.MemberMapTasksDoneMonth))
	for key, value := range jiraData.MemberMapTasksDoneMonth {
		result += fmt.Sprintf("%s выполнил(а) %d задач\n", key, value)
	}
	return result, nil
}
