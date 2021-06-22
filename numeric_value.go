package jsonutil

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// IntBool is bool value represented by int.
type IntBool bool

// Bool returns value as bool.
func (b IntBool) Bool() bool {
	return bool(b)
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *IntBool) UnmarshalJSON(data []byte) error {
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
		return fmt.Errorf("invalid int-bool value: %s", data)
	}
}

// MarshalJSON implements json.Marshaler.
func (s IntBool) MarshalJSON() ([]byte, error) {
	if s {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

// IntString is string value represented by int.
// If it is an empty string, JSON string will be `null`.
type IntString string

// String returns value as string.
func (s IntString) String() string {
	return string(s)
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *IntString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var i int
	if err := json.Unmarshal(data, &i); err != nil {
		return fmt.Errorf("invalid int-string value: %s", data)
	}

	*s = IntString(strconv.Itoa(i))
	return nil
}

// MarshalJSON implements json.Marshaler.
func (s IntString) MarshalJSON() ([]byte, error) {
	if s == "" {
		return Null, nil
	}

	i, err := strconv.Atoi(string(s))
	if err != nil {
		return nil, err
	}

	return json.Marshal(i)
}
