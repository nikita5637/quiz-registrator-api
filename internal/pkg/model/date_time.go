package model

import (
	"fmt"
	"time"
)

const (
	// TimeZoneMoscow ...
	TimeZoneMoscow = "Europe/Moscow"
)

var (
	months = []string{
		"января",
		"февраля",
		"марта",
		"апреля",
		"мая",
		"июня",
		"июля",
		"августа",
		"сентября",
		"октября",
		"ноября",
		"декабря",
	}

	days = []string{
		"ВС",
		"ПН",
		"ВТ",
		"СР",
		"ЧТ",
		"ПТ",
		"СБ",
	}
)

// DateTime ...
type DateTime time.Time

// String ...
func (d DateTime) String() string {
	if time.Time(d).IsZero() {
		return "time is empty"
	}

	loc, err := time.LoadLocation(TimeZoneMoscow)
	if err != nil {
		return ""
	}

	t := time.Time(d).In(loc)
	return fmt.Sprintf("[%s] %02d %s: %02d:%02d", days[t.Weekday()], t.Day(), months[t.Month()-1], t.Hour(), t.Minute())
}

// AsTime returns time in UTC
func (d DateTime) AsTime() time.Time {
	return time.Time(d).UTC()
}

// MarshalJSON ...
func (d DateTime) MarshalJSON() ([]byte, error) {
	return time.Time(d).MarshalJSON()
}

// UnmarshalJSON ...
func (d *DateTime) UnmarshalJSON(data []byte) error {
	t := time.Time(*d)
	if err := t.UnmarshalJSON(data); err != nil {
		return err
	}

	*d = DateTime(t)
	return nil
}
