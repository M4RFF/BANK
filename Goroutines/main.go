package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long talking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	// go greet("It's my first golang course!")
	// go greet("How is it going?")

	// creating a channel below:
	done := make(chan bool)
	go slowGreet("How ... is ... it ... going ...?", done)
	// fmt.Println(<-done)
	<-done
	// go greet("I hope you're doing great")
}
