package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	totalRoutines := 5
	wg.Add(totalRoutines)
	for i := 0; i < totalRoutines; i++ {
		go func(routineID int) {
			fmt.Printf("ID:%d - Hello goroutines!\n", routineID)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
