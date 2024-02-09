package WorkQueueServices

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	mq "scheduler/services/MQServices"
	tasklets "scheduler/tasklets"

	"go.mongodb.org/mongo-driver/bson"
)

type TaskExecution struct {
	MQInstance *mq.MQInstance
}

type TaskExecutionMetadata struct {
	TaskName string        `json:"TaskName"`
	Args     []interface{} `json:"Args"`
}

func TaskExecutionInit() *TaskExecution {
	taskExecution := &TaskExecution{}
	taskExecution.MQInstance = mq.MQInstanceInit()
	taskExecution.MQInstance.ConnectQueue("tasks")

	return taskExecution
}

func (taskExecution *TaskExecution) Enqueue(fn interface{}, args ...interface{}) (err error) {
	fnVal := reflect.ValueOf(fn)

	taskName := runtime.FuncForPC(fnVal.Pointer()).Name()
	fmt.Println("Enqueueing task: ", taskName)

	if fnVal.Kind() != reflect.Func {
		return fmt.Errorf("expected a function, got %T", fn)
	}

	taskExecutionMetadata := TaskExecutionMetadata{
		TaskName: strings.Split(taskName, ".")[1],
		Args:     args,
	}

	taskExecutionMetadataBytes, err := bson.Marshal(taskExecutionMetadata)
	if err != nil {
		return err
	}
	taskExecution.MQInstance.PublishMessage(taskExecutionMetadataBytes)
	return nil
}

func (taskExecution *TaskExecution) Execute(fnName string, args ...interface{}) (result interface{}, err error) {
	fn := tasklets.TaskRegistry[fnName]
	return fn(args...)
}
