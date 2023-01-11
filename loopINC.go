package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func add(size int, wg *sync.WaitGroup) {
	for i := 0; i < size; i++ {
		atomic.AddInt64(&counter, 1)
	}

	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	go add(1000, &wg)
	go add(1000, &wg)
	go add(1000, &wg)
	go add(1000, &wg)
	go add(1000, &wg)

	wg.Wait()
	fmt.Println(counter)
}
