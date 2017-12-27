package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println(v1)

	fmt.Println("lem of m : ", len(m))

	delete(m, "k2")
	fmt.Println("lem of m : ", len(m))

	_, prs := m["k2"]
	fmt.Println("prs : ", prs)

	_, prs = m["k1"]
	fmt.Println("check value for k1 : ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
