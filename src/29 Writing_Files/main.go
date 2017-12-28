package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	if err := ioutil.WriteFile("/tmp/dat1", d1, 0644); err != nil {
		panic(err)
	}

	file, err := os.Create("/tmp/dat2")
	check(err)

	defer file.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := file.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := file.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	file.Sync()

	w := bufio.NewWriter(file)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d butes\n", n4)

	w.Flush()
}
