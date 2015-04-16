package main

import "fmt"

func isMultipleOf(value int, isMultipleOf int) bool {
  return value % isMultipleOf == 0
}

func isMultipleOf3(value int) bool {
  return isMultipleOf(value, 3)
}

func isMultipleOf5(value int) bool {
  return isMultipleOf(value, 5)
}


func fizzbuzzOf(i int) string {
  if isMultipleOf3(i) && isMultipleOf5(i) {
    return "fizzbuzz"
  }
  if isMultipleOf3(i) {
    return "fizz"
  }
  if isMultipleOf5(i) {
    return "buzz"
  }
  return fmt.Sprintf("%d",i) // there must be a better way, returning and int as string... js you miss me.
}

func main() {
  for i:=1;i<=100;i++ {
    fmt.Println(fizzbuzzOf(i))
  }

}
