package main

//go:generate go run codegen/generate_tasklets.go

import (
	taskExecution "scheduler/services/WorkQueueServices"
	tasklets "scheduler/tasklets"
	worker "scheduler/worker"
	"time"
)

func main() {
	ch := make(chan struct{})

	go worker.Worker(ch)
	execution := taskExecution.TaskExecutionInit()
	execution.Enqueue(tasklets.IsPrime, 5)
	time.Sleep(5 * time.Second)

	execution.Enqueue(tasklets.IsPrime, 5)
}
