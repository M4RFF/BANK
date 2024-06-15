package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true // writes data to this channel

}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long talking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // writes data to this channel
	close(doneChan)
}

func main() {
	// dones := make([]chan bool, 4) // creating a slice of channels
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("It's my first golang course!", done)
	// dones[1] = make(chan bool)
	go greet("How is it going?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... is ... it ... going ...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you're doing great", done)

	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
		// fmt.Println(doneChan)
	}
}
