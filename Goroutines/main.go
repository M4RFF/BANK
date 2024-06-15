package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second) // simulate a slow, long talking task
	fmt.Println("Hello!", phrase)
}

func main() {
	go greet("It's my first golang course!")
	go greet("How is it going?")
	go slowGreet("How ... is ... it ... going ...?")
	go greet("I hope you're doing great")
}
