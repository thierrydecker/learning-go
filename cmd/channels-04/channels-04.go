package main

import (
	"fmt"
)

func send1(c1 chan string, c2 chan string, c3 chan string) {
	message := "(send1) Sent message"
	for i := 0; ; i++ {
		select {
		case c1 <- message:
			fmt.Printf("%s to %v\n", message, c1)
		case c2 <- message:
			fmt.Printf("%s to %v\n", message, c2)
		case c3 <- message:
			fmt.Printf("%s to %v\n", message, c3)
		}
	}
}

func send2(c1 chan string, c2 chan string, c3 chan string) {
	message := "(send2) Sent message"
	for i := 0; ; i++ {
		select {
		case c1 <- message:
			fmt.Printf("%s to %v\n", message, c1)
		case c2 <- message:
			fmt.Printf("%s to %v\n", message, c2)
		case c3 <- message:
			fmt.Printf("%s to %v\n", message, c3)
		}
	}
}
func receive1(c1 chan string, c2 chan string, c3 chan string) {
	for i := 0; ; i++ {
		select {
		case message := <-c1:
			fmt.Printf("(Receive1) received message (%s) from %v\n", message, c1)
		case message := <-c2:
			fmt.Printf("(Receive1) received message (%s) from %v\n", message, c2)
		case message := <-c3:
			fmt.Printf("(Receive1) received message (%s) from %v\n", message, c3)
		}
	}
}

func receive2(c1 chan string, c2 chan string, c3 chan string) {
	for i := 0; ; i++ {
		select {
		case message := <-c1:
			fmt.Printf("(Receive2) received message (%s) from %v\n", message, c1)
		case message := <-c2:
			fmt.Printf("(Receive1) received message (%s) from %v\n", message, c2)
		case message := <-c3:
			fmt.Printf("(Receive1) received message (%s) from %v\n", message, c3)
		}
	}
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	go send1(c1, c2, c3)
	go send1(c1, c2, c3)
	go send2(c2, c2, c3)
	go send2(c2, c2, c3)
	go receive1(c1, c2, c3)
	go receive1(c1, c2, c3)
	go receive2(c1, c2, c3)
	go receive2(c1, c2, c3)

	var input string
	fmt.Scanln(&input)
}
