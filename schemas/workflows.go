package schemas

type WorkflowScheduleSchema struct {
	Id                    string             `json:"id"`
	HasEditPermission     bool               `json:"hasEditPermission"`
	HasRunPermission      bool               `json:"hasRunPermission"`
	Name                  string             `json:"name"`
	Description           string             `json:"description"`
	Enabled               string             `json:"enabled"`
	CreatedBy             string             `json:"createdBy"`
	CreatedOn             string             `json:"createdOn"`
	ModifiedBy            string             `json:"modifiedBy"`
	LastModified          string             `json:"lastModified"`
	ConfigurationError    string             `json:"configurationError"`
	WorkflowInfo          WorkflowInfoSchema `json:"workflowInfo"`
	OwnerType             string             `json:"ownerType"`
	CatalogueOwner        string             `json:"catalogueOwner"`
	OwnerUserGroup        string             `json:"ownerUserGroup"`
	OwnerUser             string             `json:"ownerUser"`
	ScheduleIntervals     string             `json:"scheduleIntervals"`
	DoNotAllocateBefore   string             `json:"doNotAllocateBefore"`
	DoNotAllocateAfter    string             `json:"doNotAllocateAfter"`
	ExcludedIntervals     string             `json:"excludedIntervals"`
	AllowConcurrentJobs   string             `json:"allowConcurrentJobs"`
	JobTimeoutMinutes     string             `json:"jobTimeoutMinutes"`
	ParametersYaml        string             `json:"parametersYaml"`
	ImportParameters      string             `json:"importParameters"`
	OccurrencesStatistics string             `json:"occurrencesStatistics"`
}

type WorkflowInfoSchema struct {
	CatalogueId             string `json:"catalogueId"`
	CatalogueName           string `json:"catalogueName"`
	SeriesId                string `json:"seriesId"`
	SeriesName              string `json:"seriesName"`
	IsBoundToSpecificObject bool   `json:"isBoundToSpecificObject"`
	ObjectId                string `json:"objectId"`
	ObjectName              string `json:"objectName"`
}

type WorkflowScheduleIntervalSchema struct {
	IntervalType  string   `json:"intervalType"`
	HourlyMinutes []string `json:"hourlyMinutes"`
	DailyTimes    []string `json:"dailyTimes"`
	WeeklyDays    []string `json:"weeklyDays"`
	MonthlyDays   []string `json:"monthlyDays"`
	YearlyMonths  []string `json:"yearlyMonths"`
}

type WorkflowOccurrencesStatisticsSchema struct {
	Prev24hr    WorkflowStatsEntrySchema  `json:"prev24hr"`
	PrevWeek    WorkflowStatsEntrySchema  `json:"prevWeek"`
	PrevMonth   WorkflowStatsEntrySchema  `json:"prevMonth"`
	LastRunInfo WorkflowLastRunInfoSchema `json:"lastRunInfo"`
}

type WorkflowLastRunInfoSchema struct {
	Status            string `json:"status"`
	ScheduledTime     string `json:"scheduledTime"`
	WorkflowSessionId string `json:"workflowSessionId"`
	Messages          string `json:"messages"`
}

type WorkflowStatsEntrySchema struct {
	FailedCount   int `json:"failedCount"`
	SuceededCount int `json:"suceededCount"`
	SkippedCount  int `json:"skippedCount"`
}
