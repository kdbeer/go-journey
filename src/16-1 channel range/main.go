package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("I'm doinng task : ", i)
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
