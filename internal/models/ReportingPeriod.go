package models

type ReportingPeriod struct {
	Name string `json`
	Room string `json`

	MeetingTimes []MeetingTime `json`
}

type MeetingTime struct {
	Day string `json`
	Period string `json`
}

type ReportingPeriodMeetingTime struct {
	ReportingPeriods []ReportingPeriod `json`
}