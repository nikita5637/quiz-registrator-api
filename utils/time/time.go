package time

import "time"

var (
	// TimeNow ...
	TimeNow = time.Now
)

// ConvertTime ...
func ConvertTime(str string) time.Time {
	timeFormat := "2006-01-02 15:04"

	t, err := time.Parse(timeFormat, str)
	if err != nil {
		return time.Time{}
	}

	return t
}
