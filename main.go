package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(addNamedVariables(2, 4))
}

func add(a, b int) int {
	return a + b
}

func addNamedVariables(a int, b int) (x, y int) {
	x = a + b
	y = a - b
	return x, y
}
