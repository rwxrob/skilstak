package main

import (
	"fmt"
	"os"
)

func main() {
	var name = "Friend"

	if n := os.Getenv("USER"); len(n) > 0 {
		name = n
	}

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println("Hi, " + name + ".")
}
