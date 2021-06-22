package jsonutil

import (
	"encoding/json"
	"fmt"
)

func ExampleIntBool_UnmarshalJSON() {
	var ls []IntBool
	err := json.Unmarshal([]byte(`[1,0,null]`), &ls)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	for i, v := range ls {
		fmt.Printf("#%d: %#v\n", i, v)
	}
	// Output:
	// #0: true
	// #1: false
	// #2: false
}

func ExampleIntBool_MarshalJSON() {
	ls := []IntBool{
		true,
		false,
	}
	bs, err := json.Marshal(ls)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("%s", bs)
	// Output:
	// [1,0]
}

func ExampleIntString_UnmarshalJSON() {
	var ls []IntString
	err := json.Unmarshal([]byte(`[42,0,-1,null]`), &ls)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	for i, v := range ls {
		fmt.Printf("#%d: %#v\n", i, v)
	}
	// Output:
	// #0: "42"
	// #1: "0"
	// #2: "-1"
	// #3: ""
}

func ExampleIntString_MarshalJSON() {
	ls := []IntString{
		"42",
		"0",
		"-1",
		"",
	}
	bs, err := json.Marshal(ls)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("%s", bs)
	// Output:
	// [42,0,-1,null]
}
