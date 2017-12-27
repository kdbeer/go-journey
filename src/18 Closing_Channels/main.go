package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recieve job : ", j)
			} else {
				fmt.Println("recieved all job")
				done <- true
				return
			}

		}
	}()

	for j := 0; j < 3; j++ {
		jobs <- j
		fmt.Println("Send job")
	}

	close(jobs)
	fmt.Println("Send all job")

	<-done
}
