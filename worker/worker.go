package worker

import (
	"encoding/json"
	"fmt"
	"log"
	mq "scheduler/services/MQServices"
	taskExecution "scheduler/services/WorkQueueServices"
)

func worker() {
	mqInstance := mq.MQInstanceInit()
	mqInstance.ConnectQueue("tasks")

	mq.MQInstanceInit().ConsumeMessages(func(bytes []byte) {
		taskExecutionMetadata := taskExecution.TaskExecutionMetadata{}
		err := json.Unmarshal(bytes, &taskExecutionMetadata)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %s", err)
			return
		}

		fmt.Println("Received task: ", taskExecutionMetadata.TaskName)
	})
}
