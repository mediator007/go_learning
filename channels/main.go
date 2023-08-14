package main

import (
	strct "channels/structures"
	"fmt"
	"sync"
)

var channelSize = 10
// var ch = make(chan strct.Event, channelSize)

var ch = make(chan strct.Event)

func SendEvent(ch chan strct.Event, event strct.Event) {
	ch <- event
}

func Processor(wg *sync.WaitGroup, increment int) {
	fmt.Println("Processor", increment, " starting")
	data := strct.EventData{increment, "DefaultOperation"}
	event := strct.Event{"DefaultEvent", data}
	//try SendEvent gorutine
	// read & write to channel blocked gorutine
	SendEvent(ch, event)
	fmt.Println("Processor", increment, " send data")
	// defer is like Finally
	defer wg.Done()
}

func Consumer(wg *sync.WaitGroup, id int) {
	fmt.Println("---Consumer", id, " starting")
	// * before type - type: pointer to a sync.WaitGroup type var
	event_id := (<-ch).EventData.Id
	fmt.Println("---Consumer", id, "get result:", event_id)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	
	for i := 0; i < channelSize; i++ {
		wg.Add(1)
		// & before argument - pointer to var wg
		// if send wg, golang copy wg and send wg value to fucntion,
		//but not wg object
		go Processor(&wg, i)
		wg.Add(1)
		go Consumer(&wg, i)
	}
	fmt.Println(" ###### Cycle end ####### ")
	wg.Wait()
}
