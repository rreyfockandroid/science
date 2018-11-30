package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	in := os.Stdin
	out := os.Stdout
	io.WriteString(out, "XXX")
	_, err := io.Copy(out, in)
	if err != nil {
		logrus.Fatal(err)
	}
}
