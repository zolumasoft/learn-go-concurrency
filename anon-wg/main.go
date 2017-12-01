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
		wg.Done()
	}("Hello World!")
	wg.Wait()
}
