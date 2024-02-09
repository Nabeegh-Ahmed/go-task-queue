package worker

import (
	"fmt"
	"log"
	mq "scheduler/services/MQServices"
	taskExecution "scheduler/services/WorkQueueServices"

	"go.mongodb.org/mongo-driver/bson"
)

func Worker() {
	mqInstance := mq.MQInstanceInit()
	mqInstance.ConnectQueue("tasks")
	executionInstance := taskExecution.TaskExecution{}

	defer mqInstance.CleanUp()

	mqInstance.ConsumeMessages(func(bytes []byte) {
		taskExecutionMetadata := taskExecution.TaskExecutionMetadata{}
		err := bson.Unmarshal(bytes, &taskExecutionMetadata)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %s", err)
			return
		}

		fmt.Println("Received task: ", taskExecutionMetadata.TaskName)
		fmt.Println(executionInstance.Execute(taskExecutionMetadata.TaskName, taskExecutionMetadata.Args...))
	})
}
