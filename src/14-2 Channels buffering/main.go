package main

import "fmt"
import "time"

func main() {

	t := time.Now()
	messages := make(chan string, 5)
	messages <- updateNumber("1")
	messages <- updateNumber("2")
	messages <- updateNumber("3")
	messages <- updateNumber("4")
	messages <- updateNumber("5")

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// fmt.Println(updateNumber("1"))
	// fmt.Println(updateNumber("2"))
	// fmt.Println(updateNumber("3"))
	// fmt.Println(updateNumber("4"))
	// fmt.Println(updateNumber("5"))

	fmt.Println("Time in used", time.Since(t))
}

func updateNumber(s string) string {
	time.Sleep(10 * time.Second)
	return s
}
