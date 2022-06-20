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

func simpleLoop(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i + 1
	}
	return sum
}

func simpleSwitch(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	}
	return 0
}

func simple2dArray() [5][5]int {
	var arr = [5][5]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			arr[i][j] = i + j
		}
	}
	return arr
}

func simpleSlice() (p1 int, p2 int) {
	var arr = []int{}
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 4)
	arr = append(arr, 5)
	arr = append(arr, 6)
	var slice1 = arr[3:4]
	var slice2 = arr[:3]
	p1 = slice1[0]
	p2 = slice2[2]

	return p1, p2
}
