package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\n start")
	t := time.Now()
	fmt.Printf("\n %s\n", t.Format("2006-01-02 15:04:05"))

}
