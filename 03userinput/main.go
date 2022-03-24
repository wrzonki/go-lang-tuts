package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rating for pizza:")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)
	fmt.Printf("Type of this rating is %T", input)

}
