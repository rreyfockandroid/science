package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("This is some text!")

	b := make([]byte, 8)
	idx := 0
	for {
		n, err := r.Read(b)
		fmt.Printf("idx=%d, n = %v err = %v b = %v\n", idx, n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
		idx ++
	}
}
