package main

import (
  "fmt"
  "main/verify"
  "cube"
)


func main() {

  closure()
  
  for i := 2; i >= -2; i-- {
    num, err := verify.IsPosInt(i)
    if err != nil {
      fmt.Println("Error: ", err)
    } else {
      fmt.Println(num, " passed")
    }
  }
  
  
  var box cube.Dims
  box.SetSize(2, 4, 6)
  
  fmt.Println("area: ", box.GetArea())
  fmt.Println("volume: ", box.GetVolume())
  
  fmt.Println("width: ", box.width)


}


func closure() {

  area := func( length, width int ) int {
    return length * width
  }

  fmt.Printf("area Type %T \n", area)

  counter := func() func() int {
    num := 0
    return func() int {
      num++
      return num
    }
  } ()

  fmt.Printf("counter Type %T \n", counter)
  fmt.Println("Count: ", counter())
  fmt.Println("Count: ", counter())
}



