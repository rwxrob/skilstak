package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Answer me these questions three, 'ere the other side ye see:")

	fmt.Print("What is your name? ")
	scanner.Scan()
	name := scanner.Text()
	log.Println(`you entered name: ` + name)

	fmt.Println("What is your quest?")
	scanner.Scan()
	quest := scanner.Text()
	log.Println(`you entered quest: ` + quest)

	switch {

	case strings.HasPrefix(strings.ToLower(name), "rob"):
		fmt.Println("What is the capital of Assyria?")

	case name == "Arthur" || name == "Arthur, King of the Britons":
		fmt.Println("What is the air speed velocity of an unladen swallow?")
		fmt.Println("How do you know so much about swallows?")
		fmt.Println("Well, you have to know these things when your King you know.")

	default:
		fmt.Println("What is your favorite colour?")
	}

}
