package main

import (
	"fmt"
	"time"
)

func worker(done chan bool, s string) {
	fmt.Println("Working...")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	fmt.Println("done for : ", s)

	done <- true
}

func main() {
	t := time.Now()
	done := make(chan bool, 5)
	go worker(done, "1")
	go worker(done, "2")
	go worker(done, "3")
	go worker(done, "4")
	go worker(done, "5")

	<-done
	<-done
	<-done
	<-done
	<-done

	fmt.Println(time.Since(t))
}
