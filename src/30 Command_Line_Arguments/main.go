package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	arguWithOutProg := os.Args[2:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(arguWithOutProg)
	fmt.Println(arg)
}
