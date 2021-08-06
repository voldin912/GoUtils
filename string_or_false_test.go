package jsonutil

import (
	"encoding/json"
	"fmt"
)

func ExampleStringOrFalse_String() {
	ls := []StringOrFalse{
		{
			IsString: true,
			Value:    "foo",
		},
		{
			IsString: true,
			Value:    "",
		},
		{
			IsString: false,
			Value:    "",
		},
		{
			IsString: false,
			Value:    "foo",
		},
	}

	for i, v := range ls {
		fmt.Printf("#%d: %#v\n", i, v.String())
	}
	// Output:
	// #0: "foo"
	// #1: ""
	// #2: ""
	// #3: ""
}

func ExampleStringOrFalse_MarshalJSON() {
	ls := []StringOrFalse{
		{
			IsString: true,
			Value:    "foo",
		},
		{
			IsString: true,
			Value:    "",
		},
		{
			IsString: false,
			Value:    "",
		},
		{
			IsString: false,
			Value:    "foo",
		},
	}

	bs, err := json.Marshal(ls)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("%s", bs)
	// Output:
	// ["foo","",false,false]
}

func ExampleStringOrFalse_UnmarshalJSON() {
	var ls []StringOrFalse
	if err := json.Unmarshal([]byte(`["foo","",false]`), &ls); err != nil {
		fmt.Printf("error: %v", err)
	}

	for i, v := range ls {
		fmt.Printf("#%d: %v, %#v\n", i, v.IsString, v.Value)
	}
	// Output:
	// #0: true, "foo"
	// #1: true, ""
	// #2: false, ""
}
