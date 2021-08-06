package jsonutil

import (
	"encoding/json"
	"fmt"
)

func ExampleRawMessage_Equal() {
	var v RawMessage
	if err := json.Unmarshal([]byte(`[]`), &v); err != nil {
		fmt.Printf("error: %v", v)
	}

	fmt.Printf("value: %s\n", v)
	fmt.Printf("is empty array: %v\n", v.Equal(EmptyArray))
	// Output:
	// value: []
	// is empty array: true
}
