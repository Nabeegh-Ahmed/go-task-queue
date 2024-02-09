package WorkQueueServices

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	mq "scheduler/services/MQServices"
	tasklets "scheduler/tasklets"
)

type TaskExecution struct {
	MQInstance *mq.MQInstance
}

type TaskExecutionMetadata struct {
	TaskName string
	Args     []interface{}
}

func TaskExecutionInit() *TaskExecution {
	taskExecution := &TaskExecution{}
	taskExecution.MQInstance = mq.MQInstanceInit()
	taskExecution.MQInstance.ConnectQueue("tasks")
	return taskExecution
}

func (taskExecution *TaskExecution) Task(fn interface{}, args ...interface{}) (err error) {
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

	taskExecutionMetadataString := fmt.Sprintf("%+v", taskExecutionMetadata)

	taskExecution.MQInstance.PublishMessage(taskExecutionMetadataString)
	return nil
}

func (taskExecution *TaskExecution) Execute(fnName string, args ...interface{}) (result interface{}, err error) {
	fn := tasklets.TaskRegistry[fnName]
	return fn(args...)
}
