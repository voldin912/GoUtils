# go-jsonutil

[![Go Reference](https://pkg.go.dev/badge/github.com/kamiaka/go-jsonutil.svg)](https://pkg.go.dev/github.com/kamiaka/go-jsonutil)

JSON Utilities for Go-lang `encoding/json` package.

## Ordered JSON object.
### MarshalEntries / UnmarshalEntries

Utility function to marshal / unmarshal JSON objects into slices in the same order.

### Example

```go
// IntEntry represents integer with key.
type IntEntry struct {
	Key string
	Int int
}

// IntEntries is the slices of IntEntry.
type IntEntries []*IntEntry

// UnmarshalJSON implements `json.Unmarshaler` interface.
func (ls *IntEntries) UnmarshalJSON(bs []byte) error {
	return jsonutil.UnmarshalEntries(bs, func(key string, data []byte) error {
		entry := &IntEntry{
			Key: key,
		}
		*ls = append(*ls, entry)
		return json.Unmarshal(data, &entry.Int)
	})
}

// MarshalJSON implements `json.Marshaler` interface.
func (ls IntEntries) MarshalJSON() ([]byte, error) {
	return jsonutil.MarshalEntries(len(ls), func(i int) *jsonutil.Entry {
		return &jsonutil.Entry{
			Key:   ls[i].Key,
			Value: ls[i].Int,
		}
	})
}
```

## License

[MIT](./LICENSE)