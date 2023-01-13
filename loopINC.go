package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func add(size int) {
	wg := sync.WaitGroup{}
	for i := 0; i < size; i++ {
		atomic.AddInt64(&counter, 1)
	}

	defer wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	go func() {
		add(1000)
	}()
	go func() {
		add(1000)
	}()
	go func() {
		add(1000)
	}()
	go func() {
		add(1000)
	}()
	go func() {
		add(1000)
	}()

	wg.Wait()
	fmt.Println(counter)
}
