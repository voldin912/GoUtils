package sql

import (
	"time"
)

// DateTime value represented by `"YYYY-MM-DD hh:mm:ss"`
type DateTime time.Time

// NewDateTime returns new Date pointer.
func NewDateTime(t time.Time) *DateTime {
	v := DateTime(t)
	return &v
}

// DateTimeFormat is time layout string for DateTime.
const DateTimeFormat = `"2006-01-02 15:04:05"`

const zeroDateTimeJSON = `"0000-00-00 00:00:00"`

// String returns date-time string.
func (dt DateTime) String() string {
	t := time.Time(dt)
	if t.IsZero() {
		return t.Format("0000-00-00 00:00:00 -0700 MST")
	}
	return t.Format("2006-01-02 15:04:05 -0700 MST")
}

// Time returns specified time.Time.
func (dt DateTime) Time() time.Time {
	return time.Time(dt)
}

// UnmarshalJSON implements json.Unmarshaler.
// If `"0000-00-00"` is specified, returns zero-value Date.
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if string(data) == zeroDateTimeJSON {
		*dt = DateTime(time.Time{})
		return nil
	}

	t, err := time.Parse(DateTimeFormat, string(data))
	if err != nil {
		return err
	}

	*dt = DateTime(t)
	return nil
}

// MarshalJSON implements json.Marshaler.
// If d is zero-value, returns `"0000-00-00"`
func (dt DateTime) MarshalJSON() ([]byte, error) {
	if time.Time(dt).IsZero() {
		return []byte(zeroDateTimeJSON), nil
	}
	return []byte(time.Time(dt).Format(DateTimeFormat)), nil
}
