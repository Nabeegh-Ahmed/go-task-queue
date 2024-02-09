package main

//go:generate go run codegen/generate_tasklets.go

func main() {
	// conn, channel, queue := mq.ConnectQueue("test2")
	// mq.PublishMessage(channel, queue, "Hello, World!")
	// mq.ConsumeMessages(channel, queue)
	// defer mq.CleanUp(conn, channel, queue)

	// work_queue.EnqueueTask(tasklets.Fib, 2)

}
