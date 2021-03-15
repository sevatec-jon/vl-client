package models

type AttendanceCode struct {
	RefId string `json`
	SchoolInfoRefId string `json`
	AttendanceCode string `json`
	AttendanceType string `json`
	AttendanceStatus string `json`
	Description string `json`
	UseForDailyAttendance string `json`
	UseForPeriodAttendance string `json`
}