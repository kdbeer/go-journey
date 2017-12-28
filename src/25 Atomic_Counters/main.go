package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	optFinal := atomic.LoadUint64(&ops)
	fmt.Println("Ops :", optFinal)
}
