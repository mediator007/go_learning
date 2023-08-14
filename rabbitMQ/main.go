package main

import (
	"fmt"
	"sync"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		fmt.Println(" Rabbit ")
	}
	wg.Wait()
}
