package main

import (
	"fmt"	
)

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
			ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
			fmt.Println(message)
	}
}

func main() {

	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}