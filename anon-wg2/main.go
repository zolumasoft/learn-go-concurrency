package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Add(-1)
	}("Hello World!")
	wg.Wait()
}
