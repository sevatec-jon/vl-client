package models

type Class struct {
	RefId string `json:CourseExternalRefId`
	CourseName string `json`
	DistrictCourseID string `json`
	ClassID int `json`
	SchoolID int `json`
	CourseID int `json`
	TeacherID int `json`
	Active int `json`
	SectionDesc string `json:SectionDescr`
	Teachers []ClassTeacher `json`

	ReportingPeriodMeetingTimes []ReportingPeriodMeetingTime `json:`
}