package time

import (
	"time"
)

var (
	// TimeNow ...
	TimeNow = time.Now().UTC
)

// ConvertTime returns time.Time in UTC
func ConvertTime(str string) time.Time {
	timeFormat := "2006-01-02 15:04"

	t, err := time.Parse(timeFormat, str)
	if err != nil {
		return time.Time{}
	}

	return t
}
