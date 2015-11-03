package main

import (
  "fmt"
)

func main() {
  for i := 1; i <= 100; i++ {
    if fb := fizzbuzz(i); fb != "" {
      fmt.Println(fb)
    } else {
      fmt.Println(i)
    }
	}
}

func fizzbuzz(i int) string {
  switch {
  case i%15 == 0:
    return "FizzBuzz"
  case i%3 == 0:
    return "Fizz"
  case i%5 == 0:
    return "Buzz"
  default:
    return ""
  }
}
