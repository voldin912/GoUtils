package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// Entry represents a JSON object key / value pair.
type Entry struct {
	Key   string
	Value interface{}
}

// Entries represents list of ordered JSON object key / value pair.
type Entries []*Entry

// UnmarshalJSON implements json.Unmarshaler interface for list of ordered JSON object entries.
// It unmarshal JSON objects into slices in the same order.
func (ls *Entries) UnmarshalJSON(bs []byte) error {
	return UnmarshalEntries(bs, func(key string, data []byte) error {
		var raw json.RawMessage
		err := json.Unmarshal(data, &raw)
		if err != nil {
			return err
		}
		*ls = append(*ls, &Entry{
			Key:   key,
			Value: raw,
		})
		return nil
	})
}

// MarshalJSON implements json.Marshaler interface for list of ordered JSON object entries.
func (ls Entries) MarshalJSON() ([]byte, error) {
	return MarshalEntries(len(ls), func(i int) *Entry {
		return ls[i]
	})
}

// UnmarshalEntries is a utility function to unmarshal JSON objects into slices in the same order.
//
// e.g., unmarshal `{"a":1,"b":2,"c":3}` to `[&{Key:"a", Int:1},...]`
//   jsonutil.UnmarshalEntries([]byte(`{"a":1,"b":2,"c":3}`), func(key string, data []byte) error {
//     entry := &IntEntry{
//       Key: key,
//     }
//     entries = append(entries, entry)
//     return json.Unmarshal(data, &entry.Int)
//   })
func UnmarshalEntries(bs []byte, resolveEntry func(key string, data []byte) error) error {
	decoder := json.NewDecoder(bytes.NewReader(bs))

	uniqMap := map[string]bool{}

	token, err := decoder.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('{') {
		if token == nil {
			return nil
		}
		return fmt.Errorf("json object expected, got: %v", token)
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		if token == json.Delim('}') {
			if _, err := decoder.Token(); err != io.EOF {
				return fmt.Errorf("unexpected value after json object")
			}
			return nil
		}

		key, ok := token.(string)
		if !ok {
			return fmt.Errorf("json object key expected, got: %v", token)
		}
		if _, ok := uniqMap[key]; ok {
			return fmt.Errorf(`duplicate key: "%s"`, key)
		}
		uniqMap[key] = true

		var value json.RawMessage
		if err := decoder.Decode(&value); err != nil {
			return err
		}

		if err := resolveEntry(key, value); err != nil {
			return err
		}
	}
}

// MarshalEntries is a utility function for
//
// e.g., `[]struct{ Key string; Value int }` to `{"a":1,"b":2,"c":3}`
//   jsonutil.MarshalEntries(len(entries), func(i int) *jsonutil.Entry {
//     return &jsonutil.Entry{
//       Key:   entries[i].Key,
//       Value: entries[i].Int,
//     }
//   })
func MarshalEntries(length int, resolveEntry func(i int) *Entry) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{'{'})

	for i := 0; i < length; i++ {
		if i != 0 {
			if _, err := buf.WriteRune(','); err != nil {
				return nil, err
			}
		}

		entry := resolveEntry(i)

		bs, err := json.Marshal(entry.Key)
		if err != nil {
			return nil, err
		}
		if _, err := buf.Write(bs); err != nil {
			return nil, err
		}

		if _, err := buf.WriteRune(':'); err != nil {
			return nil, err
		}

		bs, err = json.Marshal(entry.Value)
		if err != nil {
			return nil, err
		}
		if _, err := buf.Write(bs); err != nil {
			return nil, err
		}
	}

	if _, err := buf.WriteRune('}'); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
