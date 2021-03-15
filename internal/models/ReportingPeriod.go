package models

type ReportingPeriod struct {
	Name string `json`
	Room string `json`
}

type ReportingPeriodMeetingTime struct {
	ReportingPeriods []ReportingPeriod `json`
}