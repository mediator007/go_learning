package main

import (
	"fmt"
	"sync"
)

type EventData struct {
	Id        int
	Operation string
}

type Event struct {
	Type      string
	EventData EventData
}

var ch = make(chan Event)

func SendEvent(ch chan Event, event Event) {
	ch <- event
}

func Processor(wg *sync.WaitGroup) {
	fmt.Println("---Processor start---")
	data := EventData{1, "DefaultOperation"}
	event := Event{"DefaultEvent", data}
	go SendEvent(ch, event)
	// defer is like Finally
	defer wg.Done()
}

func Consumer(wg *sync.WaitGroup) {
	fmt.Println("---Consumer start---")
	event_id := (<-ch).EventData.Id
	fmt.Println(event_id)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go Processor(&wg)
	wg.Add(1)
	go Consumer(&wg)
	wg.Wait()
}
