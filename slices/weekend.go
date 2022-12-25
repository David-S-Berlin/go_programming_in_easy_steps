package slices

import "fmt"


func OutputWeekend() {
  days := [7]string {"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
  weekend := days [5:]

  fmt.Printf("Slice weekend %v \n", weekend)
  fmt.Printf("Type weekand %T \n", weekend)
  
  fmt.Printf("Length weekend %v \n", len(weekend))
  fmt.Printf("Capacity weekend %v \n", cap(weekend))
  
  fmt.Printf("Pointer weekend: %p \n", weekend)
  fmt.Printf("Address days[5]: %p \n", &days[5])
}
