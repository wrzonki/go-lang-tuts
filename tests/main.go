package main

import (
  "fmt"
  "sort"
)

func main() {
  mySlice := []int{8, 1, 2, 3, 4, 5}
  fmt.Println(add(mySlice...))

  sortNumbers(mySlice)
  fmt.Println(mySlice)
}

func add(numbers ...int) int {
  sum := 0
  for _, num := range numbers {
    sum += num
  }
  return sum
}

//sort slice of numbers
func sortNumbers(numbers []int) {
  sort.Ints(numbers)
}
