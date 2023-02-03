package usecase

import (
	"fmt"
	"time"
	"weatherEveryDay/internal/httpClient"
	"weatherEveryDay/internal/jira"
	"weatherEveryDay/internal/models"
	"weatherEveryDay/templates"
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
			Assignees:   make(map[string]bool),
			Description: desc,
			Peredogovor: 0,
			Status:      issue.Fields.Status.Name,
		}

		changeLog, err := httpClient.FetchIssueChangelog(params.Login, params.Token, issue.Id)
		if err != nil {
			return result, err
		}

		for _, value := range changeLog.Values {
			event := value.Items[0]
			if event.Field == "duedate" && event.From != event.To {
				task.Peredogovor += 1
			}
		}

		if assignee != "" {
			task.Assignees[assignee] = true
		}

		for _, mem := range issue.Fields.Workers {
			task.Assignees[mem.DisplayName] = true
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
				for name, _ := range task.Assignees {
					result.MemberMapTasksDoneMonth[name] += 1
					result.MemberMapPeredogovorsMonth[name] += task.Peredogovor
				}
			}
			if task.DueTo.Add(time.Hour * 24 * 7).After(time.Now()) {
				result.TotalWeek += 1
				for name, _ := range task.Assignees {
					result.MemberMapTasksDoneWeek[name] += 1
					result.MemberMapPeredogovorsWeek[name] += task.Peredogovor
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
	result += fmt.Sprintf(templates.AllTask, jiraData.TotalIssues)
	result += fmt.Sprintf(templates.TaskDoneMonthWeek, jiraData.TotalMonth, jiraData.TotalWeek)
	result += fmt.Sprintf(templates.WeekMemberActivity, len(jiraData.MemberMapTasksDoneWeek))
	for key, value := range jiraData.MemberMapTasksDoneWeek {
		result += fmt.Sprintf(templates.MemberReport, key, value, jiraData.MemberMapPeredogovorsWeek[key])
	}
	result += fmt.Sprintf(templates.MonthMemberActivity, len(jiraData.MemberMapTasksDoneMonth))
	for key, value := range jiraData.MemberMapTasksDoneMonth {
		result += fmt.Sprintf(templates.MemberReport, key, value, jiraData.MemberMapPeredogovorsMonth[key])
	}
	return result, nil
}
