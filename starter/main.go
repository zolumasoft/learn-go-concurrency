package main

import (
	"fmt"
	"time"
)

func main() {
	go helloworld()
	time.Sleep(1 * time.Second)
}

func helloworld() {
	fmt.Println("Hello World!")
}
