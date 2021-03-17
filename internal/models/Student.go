package models

import "time"

type Student struct {
	FirstName string `json`
	LastName string `json`
	Gender string `json`
	MiddleName string `json:MidName`
	StudentID int `json`
	EnrollmentStatusCd string `json`
	Active int `json`
	ExternalRefId string `json`
	GradeLevelDesc string `json:GradeLevelDescr`
	SchoolID int `json`
	DateOfBirth time.Time `json`
}