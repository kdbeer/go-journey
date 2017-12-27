package main

import (
	"fmt"
	"time"
)

// We often want to execute Go code at some point in the future,
// or repeatedly at some interval. Go’s built-in timer and ticker features make both of these tasks easy.
// We’ll look first at timers and then at tickers.

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("timer1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stoped")
	}

}
