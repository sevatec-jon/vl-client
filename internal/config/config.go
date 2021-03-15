package config

type Configuration struct {
	Token string `yaml:config`
	GBDistrictID string `yaml:config`
	GBSchoolID string `yaml:config`
	OcmId string `yaml:config`
	DBConn string `yaml:config`
	SchoolId string `yaml:config`
	SchoolRefId string `yaml:config`
}