package jsonutil

import (
	"encoding/json"
	"errors"
)

// RawMessage is a raw encoded JSON value.
// It extends encoding/json.RawMessage.
type RawMessage json.RawMessage

// Equal returns whether m equals data.
func (m RawMessage) Equal(data []byte) bool {
	return string(m) == string(data)
}

// MarshalJSON returns m as the JSON encoding of m.
//
// It was copied from json.RawMessage.
// See, https://cs.opensource.google/go/go/+/refs/tags/go1.16.4:src/encoding/json/stream.go;l=263
func (m RawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return Null, nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
//
// It was copied from json.RawMessage.
// See, https://cs.opensource.google/go/go/+/refs/tags/go1.16.4:src/encoding/json/stream.go;l=271
func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("jsonutil.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
