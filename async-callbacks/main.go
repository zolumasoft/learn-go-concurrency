package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	toUpperAsync("Hello Callbacks!", func(v string) {
		fmt.Printf("Callback: %s\n", v)
		wg.Done()
	})
	fmt.Println("Awaiting async response...")
	wg.Wait()
}

func toUpperAsync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}
