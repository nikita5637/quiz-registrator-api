package model

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// TimeZoneMoscow ...
	TimeZoneMoscow = "Europe/Moscow"
)

// DateTime ...
type DateTime time.Time

// AsTime returns time
func (d DateTime) AsTime() time.Time {
	return time.Time(d)
}

// MarshalJSON ...
func (d DateTime) MarshalJSON() ([]byte, error) {
	return time.Time(d).MarshalJSON()
}

// String ...
func (d DateTime) String() string {
	return time.Time(d).String()
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

// ValidateDateTime ...
func ValidateDateTime(value interface{}) error {
	dateTime, ok := value.(DateTime)
	if !ok {
		return errors.New("must be DateTime")
	}

	return validation.Validate(dateTime.AsTime().Unix(), validation.Required, validation.Min(1))
}
