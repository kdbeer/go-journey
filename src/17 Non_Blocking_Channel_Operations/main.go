package main

import "fmt"

// Basic sends and receives on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("recieved message : ", msg)
	default:
		fmt.Println("no message recieve")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message : ", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("recieved message : ", msg)
	case sig := <-signals:
		fmt.Println("recieved signels : ", sig)
	default:
		fmt.Println("no activity")
	}
}

func sendMessage(m string, channel chan string) chan string {
	channel <- m
	return channel
}
