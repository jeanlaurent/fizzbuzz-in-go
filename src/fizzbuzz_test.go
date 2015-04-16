package main

import "testing"

func assertEquals(value string, expected string, t *testing.T) {
  if value != expected {
    t.Error("not good")
  }
}

func TestFizzbuzzOf1(t *testing.T) {
  assertEquals(fizzbuzzOf(1), "1", t)
}

func TestFizzbuzzOf3(t *testing.T) {
  assertEquals(fizzbuzzOf(3), "fizz", t)
}

func TestFizzbuzzOf5(t *testing.T) {
  assertEquals(fizzbuzzOf(5), "buzz", t)
}

func TestFizzbuzzOf15(t *testing.T) {
  assertEquals(fizzbuzzOf(15), "fizzbuzz", t)
}

// func ExamplePolka() {
//   fmt.Println("yahoo")
//   // Output:
//   // yahoo
// }
