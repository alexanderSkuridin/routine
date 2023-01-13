package main

import (
	"fmt"
	"sync"
)

var messages []string

func push(message string, wg *sync.WaitGroup) {
	mu := sync.Mutex{}
	for i := 0; i < 100; i++ {
		mu.Lock()
		messages = append(messages, message)
		mu.Unlock()
	}
	wg.Done()

}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	go push("Message", &wg)
	go push("Message", &wg)
	go push("Message", &wg)
	go push("Message", &wg)
	go push("Message", &wg)

	wg.Wait()
	fmt.Println(messages)
}
