package main

import (
	"fmt"

	"github.com/mshtlr/go_first/greetings"
)

func main() {
	message := greetings.Hello("Shootky")
	fmt.Println(message)
}
