package sql

import (
	"time"
)

// Date value represented by `"YYYY-MM-DD"`.
type Date time.Time

// DateFormat is time layout string for Date.
const DateFormat = `"2006-01-02"`

const zeroDateJSON = `"0000-00-00"`

// String returns date string.
func (d Date) String() string {
	t := time.Time(d)
	if t.IsZero() {
		return t.Format("0000-00-00 -0700 MST")
	}
	return t.Format("2006-01-02 -0700 MST")
}

// Time returns specified time.Time.
func (d Date) Time() time.Time {
	return time.Time(d)
}

// UnmarshalJSON implements json.Unmarshaler.
// If `"0000-00-00"` is specified, returns zero-value Date.
func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if string(data) == zeroDateJSON {
		*d = Date(time.Time{})
		return nil
	}

	t, err := time.Parse(DateFormat, string(data))
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

// MarshalJSON implements json.Marshaler.
// If d is zero-value, returns `"0000-00-00"`
func (d Date) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte(zeroDateJSON), nil
	}
	return []byte(time.Time(d).Format(DateFormat)), nil
}
