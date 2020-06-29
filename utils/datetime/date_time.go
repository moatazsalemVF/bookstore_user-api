package datetime

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05.000Z"
)

//GetCurrentTime returns the current time in UTC
func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

//GetCurrentTimeUTC returns a UTC time
func GetCurrentTimeUTC() string {
	time := GetCurrentTime()
	return time.Format(apiDateLayout)
}
