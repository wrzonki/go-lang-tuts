package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to mymaths")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4.6
	// fmt.Println(mynumberOne + mynumberTwo)

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRandomNumber)
}
