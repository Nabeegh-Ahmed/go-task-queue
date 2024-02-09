package main

import (
	mq "scheduler/services"
)

func main() {
	conn, channel, queue := mq.ConnectQueue("test2")
	mq.PublishMessage(channel, queue, "Hello, World!")
	mq.ConsumeMessage(channel, queue)
	defer mq.CleanUp(conn, channel, queue)
}
