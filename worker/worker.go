package worker

import "fmt"

func SomeFunc() {
	fmt.Println("Hello from worker")
}

func worker() {
	// [1] Create a connection to RabbitMQ
	// [2] Create a channel
	// [3] Declare a queue
	// [4] Create a context with a timeout
	// [5] Publish a message
}
