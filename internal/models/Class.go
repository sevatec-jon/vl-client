package models

type Class struct {
	RefId string `json:CourseExternalRefId`
	CourseName string `json`
	DistrictCourseID string `json`
	ClassID int `json`
	SchoolID int `json`
	CourseID int `json`
	Active int `json`
	SectionDesc string `json:SectionDescr`
	Teachers []ClassTeacher `json`
	Students []Student `json`
	ReportingPeriodMeetingTimes []ReportingPeriodMeetingTime `json`
}