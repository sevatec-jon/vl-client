package models

type ClassTeacher struct {
	ClassAlias string `json`
	Role string `json`
	TeacherID int `json`
}

type Teacher struct {
	FirstName string `json`
	LastName string `json`
	TeacherID int `json`
	DistrictStaffID string `json`
	Active int `json`
	ExternalRefId string `json`
}