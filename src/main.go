package main

import (
	"fmt"
)

func say(text string, c chan<- string) {
	c <- text
	fmt.Println(text)
}

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "message1"
	c <- "message2"

	fmt.Println(len(c), cap(c))

	//Close
	close(c)

	//Range
	for message := range c {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Message1", email1)
	go message("Message2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email received from email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email received from email 2", m2)
		}
	}
}
