package sql

import (
	"encoding/json"
	"fmt"
	"time"
)

func ExampleDate_MarshalJSON() {
	ls := []*Date{
		datePtr(time.Date(2021, time.July, 7, 0, 0, 0, 0, time.UTC)),
		datePtr(time.Time{}),
		nil,
	}
	bs, err := json.Marshal(ls)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%s", bs)
	// Output:
	// ["2021-07-07","0000-00-00",null]
}

func datePtr(t time.Time) *Date {
	d := Date(t)
	return &d
}

func ExampleDate_UnmarshalJSON() {
	var ls []*Date
	js := `["2021-07-07","0000-00-00",null]`
	if err := json.Unmarshal([]byte(js), &ls); err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for i, v := range ls {
		if v == nil {
			fmt.Printf("#%d: nil\n", i)
		} else {
			fmt.Printf("#%d: %v\n", i, v)
		}
	}
	// Output:
	// #0: 2021-07-07 +0000 UTC
	// #1: 0000-00-00 +0000 UTC
	// #2: nil
}
