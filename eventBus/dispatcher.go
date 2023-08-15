package main

import (
	"log"
)

func dispatcher(id int, dispatcher_channel chan string) {
	for {
		data := <- dispatcher_channel
		log.Printf("Disptcher # %d get data: %s", id, data)
	}
}



