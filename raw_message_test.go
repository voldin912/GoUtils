package jsonutil

import (
	"encoding/json"
	"fmt"
)

func ExampleRawMessage_Equal() {
	var v RawMessage
	json.Unmarshal([]byte(`[]`), &v)

	fmt.Printf("value: %s\n", v)
	fmt.Printf("is empty array: %v\n", v.Equal(EmptyArray))
	// Output:
	// value: []
	// is empty array: true
}
