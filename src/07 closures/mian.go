package main

import (
	"fmt"
)

func inseq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextint := inseq()

	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())

	newInts := inseq()
	fmt.Println(newInts())
}
