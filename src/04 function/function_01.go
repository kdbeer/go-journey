package main

import (
	"fmt"
)

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = PlusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}

func PlusPlus(a, b, c int) int {
	return a + b + c
}

func plus(a int, b int) int {
	return a + b
}
