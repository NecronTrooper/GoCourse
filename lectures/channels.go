package main

import (
	"fmt"
	"time"
) // channels used to communicate between goroutines

func main() {
	var ch chan string

	ch = make(chan string)

	go GoroutineWork(ch)

	ch <- "Hello world"

	fmt.Println("massage after writing channel")
	fmt.Scanln()

	// a <- chan int = read o nly
	//select{} = selects random, literally like switch case but for channels
	//bad practice = closing channel in place where you read
	//
}

func GoroutineWork(input chan string) {
	time.Sleep(2 * time.Second)
	inputValue, ok := <-input

	//ok status of channel
	if !ok {
		fmt.Println("channel closered")
		return
	}
	fmt.Printf("Receive massage from channel: %s\n", inputValue)
}
