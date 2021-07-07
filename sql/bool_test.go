package sql

import (
	"encoding/json"
	"fmt"
)

func ExampleBool_UnmarshalJSON() {
	var ls []Bool
	err := json.Unmarshal([]byte(`[1,0,null]`), &ls)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	for i, v := range ls {
		fmt.Printf("#%d: %t\n", i, v)
	}
	// Output:
	// #0: true
	// #1: false
	// #2: false
}

func ExampleBool_MarshalJSON() {
	ls := []Bool{
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
