package setdate

import (
	"fmt"
	"time"
)

func DateOperations() {

  // format date and time
  dateNow := time.Now()
  // Attention!!! Use everytime exactly the date-time
  // 2006-01-02 15:04:05 
  // for formatting
  dateString := dateNow.Format("2006-01-02, (Monday), 15:04:05")
  fmt.Printf("\n\nCustom format %v\n", dateString)

	// create own Date
	dt := time.Date(2025, time.January, 1, 12, 0, 0, 0, time.Local)
	fmt.Printf("\nDateTime: %v \n\n", dt)

	// AddDate
	dt = dt.AddDate(2, 6, 3)
	fmt.Printf("Added DateTime: %v \n\n", dt)

	// own layout
	layout := "2006-Jan-02 03:04PM"
	str := "2030-Dec-25 12:30AM"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Parsed Time %v \n", t)
	}

}
