package main

import (
	"fmt"
	"time"
)

func send(id string, c chan<- string) {
	for i := 0; ; i++ {
		message := fmt.Sprintf("%s - Sent message - %d", id, i)
		c <- message
		time.Sleep(time.Millisecond * 5)
	}
}

func receive(id string, c <-chan string) {
	message := ""
	for i := 0; ; i++ {
		message = <-c
		fmt.Printf("%s - Received message - %d (%s)\n", id, i, message)
	}
}

func main() {
	var c chan string = make(chan string)
	go send("Sender 1", c)
	go send("Sender 2", c)
	go send("Sender 3", c)
	go send("Sender 4", c)
	go receive("Receiver 1", c)
	go receive("Receiver 2", c)
	var input string
	fmt.Scanln(&input)
}
