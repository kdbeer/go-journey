package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(">>", n)
		}
		done <- true
	}()

	go func() {
		for j := range c {
			select {
			case n := <-c:
				fmt.Println("get something : ", n, " > ", j)
			default:
				fmt.Println("nothing to get")
			}
		}
		done <- true
	}()

	<-done
	<-done
}
