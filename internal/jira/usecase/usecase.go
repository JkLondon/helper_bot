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

func (j *JiraUC) ParseRawDataToFocusReport(params models.JiraRawData) (result models.JiraFocusData, err error) {
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

func (j *JiraUC) MakeFocusReport(params models.JiraRawData) (result string, err error) {
	jiraData, err := j.ParseRawDataToFocusReport(params)
	if err != nil {
		return err.Error(), err
	}
	result += fmt.Sprintf(templates.ReportIntro, time.Now().Format("02\\.01\\.2006"))
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

func (j *JiraUC) ParseRawDataToDailyReport(params models.JiraRawData) (result models.JiraDailyData, err error) {
	for _, issue := range params.Issues {
		if issue.Fields.Customfield10020 == nil || issue.Fields.Customfield10020[0].State != "active" {
			continue
		}
		dueTo := time.Time{}
		if issue.Fields.Duedate != nil {
			shortForm := "2006-01-02"
			dueDateStr := *issue.Fields.Duedate
			dueTo, err = time.Parse(shortForm, dueDateStr)
			if err != nil {
				return result, err
			}
		}
		dateCreation, err := time.Parse(issue.Fields.Created, "2006-01-02T15:04:05.000-0700")
		if err != nil {
			return result, err
		}
		switch issue.Fields.Issuetype.Name {
		case templates.JiraBug:
			result.FixesAllCount += 1
			switch issue.Fields.Status.Name {
			case templates.JiraStatusInProgress:
				result.FixesInProgress += 1
				if dateCreation.Add(time.Hour * 24 * 1).After(time.Now()) {
					result.FixesNewFromYesteraday += 1
				}
				if dueTo.Add(time.Hour * 24 * 1).After(time.Now()) {
					result.FixesDLToday += 1
				}
			case templates.JiraStatusDone:
				result.FixesDone += 1
				changeLog, err := httpClient.FetchIssueChangelog(params.Login, params.Token, issue.Id)
				if err != nil {
					return result, err
				}
				for _, value := range changeLog.Values {
					dateDoneChange, err := time.Parse(issue.Fields.Created, "2006-01-02sss")
					if err != nil {
						return result, err
					}
					for _, item := range value.Items {
						if item.Field == "status" && item.ToString == templates.JiraStatusDone &&
							dateDoneChange.Add(time.Hour*24*2).After(time.Now()) {
							result.FixesDoneTodayYesterday += 1
						}
					}
				}
			case templates.JiraStatusToComplete:
				result.FixesToComplete += 1
				if dueTo.Add(time.Hour * 24 * 1).After(time.Now()) {
					result.FixesDLToday += 1
				}
			}
			if issue.Fields.Customfield10021 != nil {
				result.FixesApproved += 1
			}
		case templates.JiraTask:
			if issue.Fields.Labels[0] == templates.JiraTestLabel {
				result.TestsAllCount += 1
				switch issue.Fields.Status.Name {
				case templates.JiraStatusInProgress:
					result.TestsInProgress += 1
					if dueTo.Add(time.Hour * 24 * 1).After(time.Now()) {
						result.TestsDLToday += 1
					}
				case templates.JiraStatusToComplete:
					result.TestsToComplete += 1
					if dueTo.Add(time.Hour * 24 * 1).After(time.Now()) {
						result.TestsDLToday += 1
					}
				case templates.JiraStatusDone:
					result.TestsDone += 1
					changeLog, err := httpClient.FetchIssueChangelog(params.Login, params.Token, issue.Id)
					if err != nil {
						return result, err
					}
					for _, value := range changeLog.Values {
						dateDoneChange, err := time.Parse(issue.Fields.Created, time.RFC3339)
						if err != nil {
							return result, err
						}
						for _, item := range value.Items {
							if item.Field == "status" && item.ToString == templates.JiraStatusDone &&
								dateDoneChange.Add(time.Hour*24*2).After(time.Now()) {
								result.TestsDoneTodayYesterday += 1
							}
						}
					}
				}
			}
		case templates.JiraIter:
			result.ItersAllCount += 1
			switch issue.Fields.Status.Name {
			case templates.JiraStatusToComplete:
				result.ItersToComplete += 1
				if dueTo.Add(time.Hour * 24 * 1).After(time.Now()) {
					result.ItersDLToday += 1
				}
			case templates.JiraStatusInProgress:
				result.ItersInProgress += 1
				if dueTo.Add(time.Hour * 24 * 1).After(time.Now()) {
					result.ItersDLToday += 1
				}
			case templates.JiraStatusToCheck:
				result.ItersToCheck += 1
				if dueTo.Add(time.Hour * 24 * 1).After(time.Now()) {
					result.ItersDLToday += 1
				}
			case templates.JiraStatusDone:
				result.ItersDone += 1
				changeLog, err := httpClient.FetchIssueChangelog(params.Login, params.Token, issue.Id)
				if err != nil {
					return result, err
				}
				for _, value := range changeLog.Values {
					dateDoneChange, err := time.Parse(issue.Fields.Created, "2006-01-02T15:04:05.000-0700")
					if err != nil {
						return result, err
					}
					for _, item := range value.Items {
						if item.Field == "status" && item.ToString == templates.JiraStatusDone &&
							dateDoneChange.Add(time.Hour*24*2).After(time.Now()) {
							result.ItersDoneTodayYesterday += 1
						}
					}
				}
			}
		}
	}
	return result, nil
}

func (j *JiraUC) MakeDailyReport(params models.JiraRawData) (result string, err error) {
	jiraData, err := j.ParseRawDataToDailyReport(params)
	if err != nil {
		return err.Error(), err
	}
	result += fmt.Sprintf(templates.EverydayReport,
		time.Now().Format("02\\.01\\.2006"),
		jiraData.FixesDLToday,
		jiraData.FixesInProgress,
		jiraData.FixesNewFromYesteraday,
		jiraData.FixesToComplete,
		jiraData.FixesDoneTodayYesterday,
		jiraData.FixesDone,
		jiraData.FixesApproved,
		jiraData.FixesAllCount,

		jiraData.TestsDLToday,
		jiraData.TestsInProgress,
		jiraData.TestsToComplete,
		jiraData.TestsAllCount,
		jiraData.TestsDoneTodayYesterday,
		jiraData.TestsDone,

		jiraData.ItersDLToday,
		jiraData.ItersInProgress,
		jiraData.ItersToComplete,
		jiraData.ItersToCheck,
		jiraData.ItersDoneTodayYesterday,
		jiraData.ItersDone,
		jiraData.ItersAllCount,
	)
	return result, nil
}
