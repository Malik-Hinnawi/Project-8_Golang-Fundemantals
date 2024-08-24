package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	// Adding go, makes it parallel
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How... are... you ...?", done)
	go greet("I hope you're liking the course!", done)

	// Solution Type 1
	// Problem: Race conditons
	for range done {
	}

	// Solution type 2
	dones := make([]chan bool, 4)
	for i := range dones {
		dones[i] = make(chan bool)
	}

	go greet("Nice to meet you!", dones[0])
	go greet("How are you?", dones[1])
	go slowGreet("How... are... you ...?", dones[2])
	go greet("I hope you're liking the course!", dones[3])

	for _, v := range dones {
		<-v
	}
}
