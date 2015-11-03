package main

import "testing"

func Test_Fizz(t *testing.T) {
  out := fizzbuzz(3)
  if out != "Fizz" {
    t.Log("fizz expected but not found")
    t.Fail()
  }
}

func Test_Buzz(t *testing.T) {
  out := fizzbuzz(5)
  if out != "Buzz" {
    t.Log("fizz expected but not found")
    t.Fail()
  }
}

func Test_FizzBuzz(t *testing.T) {
  out := fizzbuzz(15)
  if out != "FizzBuzz" {
    t.Log("fizz expected but not found")
    t.Fail()
  }
}
