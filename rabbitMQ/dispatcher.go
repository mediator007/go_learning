package main

import (
	"fmt"
)

func dispatcher(id int) {
	var forever chan struct{}
	fmt.Printf("Disptcher # %d get data", id)
	<-forever
}