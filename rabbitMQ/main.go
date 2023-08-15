package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 10; i++ {
		go send(i)
		go recieve(i)
	}
	wg.Wait()
}
