package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("foo.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	fmt.Println(string(buf[:n]))
}
