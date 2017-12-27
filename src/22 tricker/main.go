package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Microsecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Ticker at : ", t)
		}
	}()

	time.Sleep(time.Second * 4)
	ticker.Stop()
	fmt.Println("Ticer stoped")
}
