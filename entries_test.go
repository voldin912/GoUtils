package jsonutil_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/kamiaka/go-jsonutil"
)

func TestEntries_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		json      string
		want      interface{}
		wantError interface{}
	}{
		{
			json: `{"a":42,"b":[1,2,3]}`,
			want: jsonutil.Entries{
				&jsonutil.Entry{
					Key:   "a",
					Value: json.RawMessage("42"),
				},
				&jsonutil.Entry{
					Key:   "b",
					Value: json.RawMessage("[1,2,3]"),
				},
			},
		},
		{
			json:      ``,
			wantError: io.EOF,
		},
		{
			json:      `[]`,
			wantError: errors.New("json object expected, got: ["),
		},
	}

	for i, tc := range cases {
		var entries jsonutil.Entries
		err := entries.UnmarshalJSON([]byte(tc.json))
		if err != nil {
			if reflect.DeepEqual(tc.wantError, err) {
				continue
			}
			t.Fatalf("#%d: Entries.UnmarshalJSON(...) returns error: %v", i, err)
		}
		if !reflect.DeepEqual(tc.want, entries) {
			t.Errorf("#%d: Entries.UnmarshalJSON(...) == %#v, got: %#v", i, entries, tc.want)
		}
	}
}

func TestEntries_MarshalJSON(t *testing.T) {
	cases := []struct {
		entries   jsonutil.Entries
		want      []byte
		wantError interface{}
	}{
		{
			entries: jsonutil.Entries{
				&jsonutil.Entry{
					Key:   "a",
					Value: json.RawMessage("42"),
				},
				&jsonutil.Entry{
					Key:   "b",
					Value: json.RawMessage("[1,2,3]"),
				},
			},
			want: []byte(`{"a":42,"b":[1,2,3]}`),
		},
	}

	for i, tc := range cases {
		got, err := tc.entries.MarshalJSON()
		if err != nil {
			if reflect.DeepEqual(tc.wantError, err) {
				continue
			}
			t.Fatalf("#%d: Entries.MarshalJSON(...) returns error: %v", i, err)
		}
		if !bytes.Equal(tc.want, got) {
			t.Errorf("#%d: Entries.MarshalJSON(...) == %s, want: %s", i, got, tc.want)
		}
	}
}

type IntEntry struct {
	Key string
	Int int
}

func ExampleUnmarshalEntries() {
	// type IntEntry struct {
	//   Key string
	//   Int  int
	// }

	var entries []*IntEntry
	if err := jsonutil.UnmarshalEntries([]byte(`{"a":1,"b":2,"c":3}`), func(key string, data []byte) error {
		entry := &IntEntry{
			Key: key,
		}
		entries = append(entries, entry)
		return json.Unmarshal(data, &entry.Int)
	}); err != nil {
		fmt.Printf("error: %v", err)
	}

	for i, v := range entries {
		fmt.Printf("#%d: %+v\n", i, v)
	}
	// Output:
	// #0: &{Key:a Int:1}
	// #1: &{Key:b Int:2}
	// #2: &{Key:c Int:3}
}

func ExampleMarshalEntries() {
	// type IntEntry struct {
	//   Key string
	//   Int  int
	// }

	entries := []*IntEntry{
		{
			Key: "a",
			Int: 42,
		},
		{
			Key: "b",
			Int: 99,
		},
	}

	bs, _ := jsonutil.MarshalEntries(len(entries), func(i int) *jsonutil.Entry {
		return &jsonutil.Entry{
			Key:   entries[i].Key,
			Value: entries[i].Int,
		}
	})
	fmt.Println(string(bs))
	// Output:
	// {"a":42,"b":99}
}

type MultiFieldEntry struct {
	Key string `json:"-"`
	A   string `json:"a"`
	B   int    `json:"b"`
}

func ExampleUnmarshalEntries_multiFields() {
	// type MultiFieldEntry struct {
	//   Key string `json:"-"`
	//   A   string `json:"a"`
	//   B   int    `json:"b"`
	// }

	var entries []*MultiFieldEntry
	if err := jsonutil.UnmarshalEntries(
		[]byte(`{"foo":{"a":"A1","b":42},"bar":{"a":"A2","b":99},"buz":null}`),
		func(key string, data []byte) error {
			var entry *MultiFieldEntry
			if err := json.Unmarshal(data, &entry); err != nil || entry == nil {
				return err
			}
			entry.Key = key
			entries = append(entries, entry)
			return nil
		},
	); err != nil {
		fmt.Printf("error: %v", err)
	}

	for i, v := range entries {
		fmt.Printf("#%d: %+v\n", i, v)
	}
	// Output:
	// #0: &{Key:foo A:A1 B:42}
	// #1: &{Key:bar A:A2 B:99}
}

func ExampleMarshalEntries_multiFields() {
	// type IntEntry struct {
	//   Key string
	//   Int  int
	// }

	entries := []*MultiFieldEntry{
		{
			Key: "foo",
			A:   "A1",
			B:   42,
		},
		{
			Key: "bar",
			A:   "A2",
			B:   99,
		},
	}

	bs, _ := jsonutil.MarshalEntries(len(entries), func(i int) *jsonutil.Entry {
		return &jsonutil.Entry{
			Key:   entries[i].Key,
			Value: entries[i],
		}
	})
	fmt.Println(string(bs))
	// Output:
	// {"foo":{"a":"A1","b":42},"bar":{"a":"A2","b":99}}
}
