package main

import (
	"fmt"
)

func zeroVal(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initail i : ", i)

	zeroVal(i)
	fmt.Println("Zeroval : ", i)

	zeroptr(&i)
	fmt.Println("zeroptr : ", i)

	fmt.Println("pointer i : ", i)

}
