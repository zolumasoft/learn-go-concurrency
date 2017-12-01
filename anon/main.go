package main

import (
	"fmt"
	"time"
)

func main() {
	go func(msg string) {
		fmt.Println(msg)
	}("Hello World!")
	time.Sleep(1 * time.Second)
}
