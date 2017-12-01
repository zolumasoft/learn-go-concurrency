package main

import (
	"fmt"
	"time"
)

func main() {
	messagePrinter := func(msg string) {
		fmt.Println(msg)
	}
	go messagePrinter("Hello, World!")
	go messagePrinter("Hello Go Routine")
	time.Sleep(1 * time.Second)
}
