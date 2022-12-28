package setdate

import (
  "fmt"
  "time"
)


func DateOperations() {

  // create own Date
  dt := time.Date(2025, time.January, 1, 12, 0, 0, 0, time.Local)
  fmt.Printf("\nDateTime: %v \n\n", dt)
  
  // AddDate
  dt = dt.AddDate(2, 6, 3)
  fmt.Printf("Added DateTime: %v \n\n", dt)

}
