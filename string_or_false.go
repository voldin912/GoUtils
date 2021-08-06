package jsonutil

import "encoding/json"

// StringOrFalse is JSON value that represents string or `false`.
type StringOrFalse struct {
	IsString bool
	Value    string
}

func (s StringOrFalse) String() string {
	if !s.IsString {
		return ""
	}
	return s.Value
}

// MarshalJSON implements encoding/json.Marshaler.
func (s StringOrFalse) MarshalJSON() ([]byte, error) {
	if !s.IsString {
		return False, nil
	}

	return json.Marshal(s.Value)
}

// UnmarshalJSON implements encoding/json.Unmarshaler.
func (s *StringOrFalse) UnmarshalJSON(data []byte) error {
	if False.Equal(data) {
		*s = StringOrFalse{
			IsString: false,
			Value:    "",
		}
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	*s = StringOrFalse{
		IsString: true,
		Value:    str,
	}

	return nil
}
