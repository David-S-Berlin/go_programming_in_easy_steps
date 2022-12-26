package main

import (
  "fmt"
  "main/verify"
  "main/cube"
  "main/slices"
)


func main() {

  closure()
  
  //verify
  for i := 2; i >= -2; i-- {
    num, err := verify.IsPosInt(i)
    if err != nil {
      fmt.Println("Error: ", err)
    } else {
      fmt.Println(num, " passed")
    }
  }
  
  // cube
  var box cube.Dims
  box.SetSize(2, 4, 6)
  fmt.Println("area: ", box.GetArea())
  fmt.Println("volume: ", box.GetVolume())

  // slices
  slices.OutputWeekend()
}




