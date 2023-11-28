package main

import (
	"fmt"
	"io"
	"os"
)

func read(r io.Reader) []byte {
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)
	return buf[:n]
}

func main() {
	f, err := os.Open("foo.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println(string(read(f)))
}
