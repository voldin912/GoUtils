package sql

import "fmt"

// Bool is bool value represented by TinyInt(1).
type Bool bool

// Bool returns value as bool.
func (b Bool) Bool() bool {
	return bool(b)
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *Bool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "1":
		*b = true
		return nil
	case "0":
		*b = false
		return nil
	case "null":
		return nil
	default:
		return fmt.Errorf("invalid bool value: %s", data)
	}
}

// MarshalJSON implements json.Marshaler.
func (s Bool) MarshalJSON() ([]byte, error) {
	if s {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}
