package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["Go"] = "Golang"
	languages["Python"] = "Python"
	languages["Java"] = "Java"
	languages["C++"] = "C++"

	fmt.Println(languages)
	fmt.Println(languages["Go"])

	delete(languages, "Java")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println("Key:", key, "Value:", value)
	}
}
