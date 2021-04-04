package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	iterations := 100

	workerNames := []string{"apple", "bravo", "charlie", "doggy", "echo"}
	for _, name := range workerNames {
		go waitGroupWorker(name, ch, &wg)
	}

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		ch <- i
	}

	close(ch)
	wg.Wait()
}

func waitGroupWorker(name string, ch chan int, wg *sync.WaitGroup) {
	for value := range ch {
		fmt.Printf("%d processed by %s\n", value, name)
		wg.Done()
	}
}
