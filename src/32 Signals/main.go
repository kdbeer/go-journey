package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sign := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sign
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awating sygnal")
	<-done
	fmt.Println("existing")
}
