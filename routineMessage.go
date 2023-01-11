package main

import (
	"fmt"
	"sync"
)

type Message interface {
	Push(string)
	Pop() string
}

func produce(m Message, mu *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		m.Push("Message")
		mu.Unlock()
	}
}
func consume(m Message, mu *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		fmt.Println(m.Pop())
		mu.Unlock()
	}
}

func main() {
	var m Message
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		produce(m, &mu)
	}()
	go func() {
		defer wg.Done()
		consume(m, &mu)
	}()

	wg.Wait()

}
