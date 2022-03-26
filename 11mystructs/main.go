package main

import "fmt"

func main() {
	pawel := User{"Pawel", "abc@gmail.com", true, 30}
	fmt.Println(pawel)
	fmt.Printf("details are: %+v\n", pawel)
	fmt.Printf("name is %v and email is %+v\n", pawel.Name, pawel.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
