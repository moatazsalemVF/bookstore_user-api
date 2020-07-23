package datetime

import "time"

const (
	apiDateLayout       = "2006-01-02T15:04:05.000Z"
	mysqlDateTimeLayout = "2006-01-02 15:04:05"
)

//GetCurrentTime returns the current time in UTC
func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

//GetCurrentTimeUTC returns a UTC time
func GetCurrentTimeUTC() string {
	return GetCurrentTime().Format(apiDateLayout)
}

//GetMysqlCurrentTimeUTC returns a UTC time in mysql format
func GetMysqlCurrentTimeUTC() string {
	return GetCurrentTime().Format(mysqlDateTimeLayout)
}
