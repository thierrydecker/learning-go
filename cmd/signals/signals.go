package main

import (
	"os"
	"os/signal"
	"fmt"
	"time"
)

func signalHandler(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	defer func() {
		fmt.Printf("Program stopped\n")
		os.Exit(0)
	}()
	for s := <-c; ; s = <-c {
		switch s {
		default:
			fmt.Printf("Received %v signal\n", s)
			return
		}
	}
}

func main() {
	fmt.Printf("Program starts an infinite loop\n")
	fmt.Printf("Waiting for any signal to stop...\n")
	go signalHandler(make(chan os.Signal))
	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 1000)
	}
}
