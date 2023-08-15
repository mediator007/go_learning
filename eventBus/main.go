package main

import (
	"sync"
)

func main() {
	var ch1 = make(chan string)
	var ch2 = make(chan string)
	var ch3 = make(chan string)

	go dispatcher(1, ch1)
	go dispatcher(2, ch2)
	go dispatcher(3, ch3)
	
	var dispatchersChannelsList []chan string
	dispatchersChannelsList = append(dispatchersChannelsList, ch1)
	dispatchersChannelsList = append(dispatchersChannelsList, ch2)
	dispatchersChannelsList = append(dispatchersChannelsList, ch3)

	var wg sync.WaitGroup
	wg.Add(1)
	
	for i := 0; i < 3; i++ {
		go sender(i)
	}
	
	reciever(1, dispatchersChannelsList)
	wg.Wait()
}
