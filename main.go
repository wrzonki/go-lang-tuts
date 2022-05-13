package main

import (
	"fmt"
	"os"
)

type Testing struct {
	A int
	B int
	C int
	D int
	E int
}

var tests = Testing{
	A: 7,
	B: 7,
	C: -1,
	D: 3,
	E: 0,
}

func main() {
	os.Exit(runMain(tests))
}

func runMain(t Testing) int {
	fmt.Println(tests)
	return 0
}

func add(a, b int) int {
	return a + b
}

func addNamedVariables(a int, b int) (x, y int) {
	x = a + b
	y = a - b
	return x, y
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}
