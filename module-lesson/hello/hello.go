package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Shawn")
	fmt.Println(message)
}
