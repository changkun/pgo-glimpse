package main

import (
	"fmt"
	"io"
	"os"
)

func read(r io.Reader) (n []byte) {
	buf := make([]byte, 1024)
	if f, ok := r.(*os.File); ok {
		n, _ = f.Read(buf)
	} else {
		n, _ = f.Read(buf)
	}
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
