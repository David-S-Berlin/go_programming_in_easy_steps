package main

import "fmt"


func main() {

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
