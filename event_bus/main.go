package main

import (
	"fmt"
	"sync"
    strct "eventbus/structures"
)

var ch = make(chan strct.Event)

func SendEvent(ch chan strct.Event, event strct.Event) {
	ch <- event
}

func Processor(wg *sync.WaitGroup) {
	fmt.Println("---Processor start---")
	data := strct.EventData{1, "DefaultOperation"}
	event := strct.Event{"DefaultEvent", data}
	go SendEvent(ch, event)
	// defer is like Finally
	defer wg.Done()
}

func Consumer(wg *sync.WaitGroup) {
    // * before type - type: pointer to a sync.WaitGroup type var
	fmt.Println("---Consumer start---")
	event_id := (<-ch).EventData.Id
	fmt.Println(event_id)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
    fmt.Println(wg)
	wg.Add(1)
    // & before argument - pointer to var wg
    // if send wg, golang copy wg and send wg value to fucntion,
    //but not wg object
	go Processor(&wg)
	wg.Add(1)
	go Consumer(&wg)
	wg.Wait()
}
