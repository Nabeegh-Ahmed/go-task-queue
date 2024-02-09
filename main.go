package main

//go:generate go run codegen/generate_tasklets.go

import (
	taskExecution "scheduler/services/WorkQueueServices"
	worker "scheduler/worker"
)

func main() {
	go worker.Worker()
	execution := taskExecution.TaskExecutionInit()
}
