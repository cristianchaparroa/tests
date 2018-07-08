package main

import (
  "fmt"
)

func thousands(number int, expression string) string {

  if (number % 1000) != 0 {
    mod := number % 1000
    pe  := number/1000

    expression = fmt.Sprintf("%v,%v",mod, expression)
    return thousands(pe, expression)
  }
  return expression[:len(expression)-1]
}

func Thousands(number int) string {
  return thousands(number,"")
}
