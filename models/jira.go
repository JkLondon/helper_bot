package models

import "time"

type JiraReport struct {
	Expand     string  `json:"expand"`
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}

type Issue struct {
	Expand string `json:"expand"`
	Id     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields Field  `json:"fields"`
}

type Field struct {
	Statuscategorychangedate string    `json:"statuscategorychangedate"`
	Issuetype                IssueType `json:"issuetype"`
	Parent                   `json:"parent,omitempty"`
	Timespent                interface{} `json:"timespent"`
	Customfield10030         interface{} `json:"customfield_10030"`
	Project                  struct {
		Self           string `json:"self"`
		Id             string `json:"id"`
		Key            string `json:"key"`
		Name           string `json:"name"`
		ProjectTypeKey string `json:"projectTypeKey"`
		Simplified     bool   `json:"simplified"`
		AvatarUrls     struct {
			X48 string `json:"48x48"`
			X24 string `json:"24x24"`
			X16 string `json:"16x16"`
			X32 string `json:"32x32"`
		} `json:"avatarUrls"`
	} `json:"project"`
	Customfield10031   interface{}   `json:"customfield_10031"`
	FixVersions        []interface{} `json:"fixVersions"`
	Aggregatetimespent interface{}   `json:"aggregatetimespent"`
	Resolution         *struct {
		Self        string `json:"self"`
		Id          string `json:"id"`
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"resolution"`
	Customfield10027 interface{} `json:"customfield_10027"`
	Customfield10028 interface{} `json:"customfield_10028"`
	Customfield10029 interface{} `json:"customfield_10029"`
	Resolutiondate   *string     `json:"resolutiondate"`
	Workratio        int         `json:"workratio"`
	Watches          struct {
		Self       string `json:"self"`
		WatchCount int    `json:"watchCount"`
		IsWatching bool   `json:"isWatching"`
	} `json:"watches"`
	LastViewed       *string     `json:"lastViewed"`
	Customfield10061 interface{} `json:"customfield_10061"`
	Created          string      `json:"created"`
	Customfield10020 []struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		State     string    `json:"state"`
		BoardId   int       `json:"boardId"`
		Goal      string    `json:"goal,omitempty"`
		StartDate time.Time `json:"startDate"`
		EndDate   time.Time `json:"endDate"`
	} `json:"customfield_10020"`
	Customfield10021 interface{} `json:"customfield_10021"`
	Customfield10022 interface{} `json:"customfield_10022"`
	Customfield10023 interface{} `json:"customfield_10023"`
	Priority         struct {
		Self    string `json:"self"`
		IconUrl string `json:"iconUrl"`
		Name    string `json:"name"`
		Id      string `json:"id"`
	} `json:"priority"`
	Customfield10024 *string     `json:"customfield_10024"`
	Customfield10025 *string     `json:"customfield_10025"`
	Customfield10026 interface{} `json:"customfield_10026"`
	Labels           []string    `json:"labels"`
	Customfield10016 *float64    `json:"customfield_10016"`
	Customfield10017 *string     `json:"customfield_10017"`
	Customfield10018 struct {
		HasEpicLinkFieldDependency bool `json:"hasEpicLinkFieldDependency"`
		ShowField                  bool `json:"showField"`
		NonEditableReason          struct {
			Reason  string `json:"reason"`
			Message string `json:"message"`
		} `json:"nonEditableReason"`
	} `json:"customfield_10018"`
	Customfield10019              string        `json:"customfield_10019"`
	Timeestimate                  interface{}   `json:"timeestimate"`
	Aggregatetimeoriginalestimate interface{}   `json:"aggregatetimeoriginalestimate"`
	Versions                      []interface{} `json:"versions"`
	Issuelinks                    []interface{} `json:"issuelinks"`
	Assignee                      *Worker       `json:"assignee"`
	Updated                       string        `json:"updated"`
	Status                        struct {
		Self           string `json:"self"`
		Description    string `json:"description"`
		IconUrl        string `json:"iconUrl"`
		Name           string `json:"name"`
		Id             string `json:"id"`
		StatusCategory struct {
			Self      string `json:"self"`
			Id        int    `json:"id"`
			Key       string `json:"key"`
			ColorName string `json:"colorName"`
			Name      string `json:"name"`
		} `json:"statusCategory"`
	} `json:"status"`
	Components            []interface{} `json:"components"`
	Timeoriginalestimate  interface{}   `json:"timeoriginalestimate"`
	Description           *string       `json:"description"`
	Customfield10010      interface{}   `json:"customfield_10010"`
	Customfield10014      interface{}   `json:"customfield_10014"`
	Customfield10015      *string       `json:"customfield_10015"`
	Customfield10005      interface{}   `json:"customfield_10005"`
	Customfield10006      interface{}   `json:"customfield_10006"`
	Customfield10007      interface{}   `json:"customfield_10007"`
	Security              interface{}   `json:"security"`
	Customfield10008      interface{}   `json:"customfield_10008"`
	Customfield10009      interface{}   `json:"customfield_10009"`
	Aggregatetimeestimate interface{}   `json:"aggregatetimeestimate"`
	Summary               string        `json:"summary"`
	Creator               Worker        `json:"creator"`
	Subtasks              []struct {
		Id     string `json:"id"`
		Key    string `json:"key"`
		Self   string `json:"self"`
		Fields struct {
			Summary string `json:"summary"`
			Status  struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconUrl        string `json:"iconUrl"`
				Name           string `json:"name"`
				Id             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					Id        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
			Priority struct {
				Self    string `json:"self"`
				IconUrl string `json:"iconUrl"`
				Name    string `json:"name"`
				Id      string `json:"id"`
			} `json:"priority"`
			Issuetype struct {
				Self           string `json:"self"`
				Id             string `json:"id"`
				Description    string `json:"description"`
				IconUrl        string `json:"iconUrl"`
				Name           string `json:"name"`
				Subtask        bool   `json:"subtask"`
				AvatarId       int    `json:"avatarId"`
				EntityId       string `json:"entityId"`
				HierarchyLevel int    `json:"hierarchyLevel"`
			} `json:"issuetype"`
		} `json:"fields"`
	} `json:"subtasks"`
	Customfield10041  interface{} `json:"customfield_10041"`
	Customfield10042  interface{} `json:"customfield_10042"`
	Customfield10043  interface{} `json:"customfield_10043"`
	Reporter          Worker      `json:"reporter"`
	Workers           []Worker    `json:"customfield_10044,omitempty"` //исполнители
	Aggregateprogress struct {
		Progress int `json:"progress"`
		Total    int `json:"total"`
	} `json:"aggregateprogress"`
	Customfield10001 interface{} `json:"customfield_10001"`
	Customfield10002 interface{} `json:"customfield_10002"`
	Customfield10003 interface{} `json:"customfield_10003"`
	Customfield10004 interface{} `json:"customfield_10004"`
	Environment      interface{} `json:"environment"`
	Duedate          *string     `json:"duedate"`
	Progress         struct {
		Progress int `json:"progress"`
		Total    int `json:"total"`
	} `json:"progress"`
	Votes struct {
		Self     string `json:"self"`
		Votes    int    `json:"votes"`
		HasVoted bool   `json:"hasVoted"`
	} `json:"votes"`
	Customfield10035 *string     `json:"customfield_10035,omitempty"`
	Customfield10062 interface{} `json:"customfield_10062"`
	Customfield10036 interface{} `json:"customfield_10036"`
	Customfield10063 interface{} `json:"customfield_10063"`
}

type IssueType struct {
	Self           string `json:"self"`
	Id             string `json:"id"`
	Description    string `json:"description"`
	IconUrl        string `json:"iconUrl"`
	Name           string `json:"name"`
	Subtask        bool   `json:"subtask"`
	AvatarId       int    `json:"avatarId"`
	EntityId       string `json:"entityId"`
	HierarchyLevel int    `json:"hierarchyLevel"`
}

type Parent struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	Self   string `json:"self"`
	Fields struct {
		Summary string `json:"summary"`
		Status  struct {
			Self           string `json:"self"`
			Description    string `json:"description"`
			IconUrl        string `json:"iconUrl"`
			Name           string `json:"name"`
			Id             string `json:"id"`
			StatusCategory struct {
				Self      string `json:"self"`
				Id        int    `json:"id"`
				Key       string `json:"key"`
				ColorName string `json:"colorName"`
				Name      string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
		Priority struct {
			Self    string `json:"self"`
			IconUrl string `json:"iconUrl"`
			Name    string `json:"name"`
			Id      string `json:"id"`
		} `json:"priority"`
		Issuetype struct {
			Self           string `json:"self"`
			Id             string `json:"id"`
			Description    string `json:"description"`
			IconUrl        string `json:"iconUrl"`
			Name           string `json:"name"`
			Subtask        bool   `json:"subtask"`
			AvatarId       int    `json:"avatarId"`
			EntityId       string `json:"entityId"`
			HierarchyLevel int    `json:"hierarchyLevel"`
		} `json:"issuetype"`
	} `json:"fields"`
}

type Worker struct {
	Self       string `json:"self"`
	AccountId  string `json:"accountId"`
	AvatarUrls struct {
		X48 string `json:"48x48"`
		X24 string `json:"24x24"`
		X16 string `json:"16x16"`
		X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	TimeZone     string `json:"timeZone"`
	AccountType  string `json:"accountType"`
	EmailAddress string `json:"emailAddress,omitempty"`
}
