package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Orange", "Banana", "Grape", "Plum"}

	fruitList = append(fruitList, "Cherry", "Kiwi", "Eggplant", "Peach")
	fmt.Println("Fruit List:", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 935
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"Go", "Python", "Java", "C++", "C#", "Ruby", "PHP", "JavaScript"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
