package sql

import (
	"encoding/json"
	"fmt"
	"time"
)

func ExampleDateTime_MarshalJSON() {
	ls := []*DateTime{
		dateTimePtr(time.Date(2021, time.July, 7, 12, 34, 56, 0, time.UTC)),
		dateTimePtr(time.Time{}),
		nil,
	}
	bs, err := json.Marshal(ls)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%s", bs)
	// Output:
	// ["2021-07-07 12:34:56","0000-00-00 00:00:00",null]
}

func dateTimePtr(t time.Time) *DateTime {
	dt := DateTime(t)
	return &dt
}

func ExampleDateTime_UnmarshalJSON() {
	var ls []*DateTime
	js := `["2021-07-07 12:34:56","0000-00-00 00:00:00",null]`
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
	// #0: 2021-07-07 12:34:56 +0000 UTC
	// #1: 0000-00-00 00:00:00 +0000 UTC
	// #2: nil
}
