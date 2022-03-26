package main

import "fmt"

func main() {
	fmt.Println("arrays")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	// fruitList[2] = "Banana"
	fruitList[3] = "Grape"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList)) // 4

	var vegList = [3]string{"Carrot", "Potato", "Cucumber"}
	fmt.Println("Veg List:", vegList)
	fmt.Println("Veg List:", len(vegList))
}
