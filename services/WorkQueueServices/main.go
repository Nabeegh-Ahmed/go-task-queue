package WorkQueueServices

import (
	"fmt"
	"reflect"
	"runtime"

	tasklets "scheduler/tasklets"
)

func EnqueueTask(fn interface{}, args ...interface{}) (result []reflect.Value, err error) {
	fnVal := reflect.ValueOf(fn)

	taskName := runtime.FuncForPC(fnVal.Pointer()).Name()
	fmt.Println("Enqueueing task: ", taskName)

	if fnVal.Kind() != reflect.Func {
		return nil, fmt.Errorf("expected a function, got %T", fn)
	}
	// TODO send this to mq
	// Prepare arguments for reflection call
	return fnVal.Call([]reflect.Value{}), nil
}

func ExecuteTask(fnName string, args ...interface{}) (result interface{}, err error) {
	fn := tasklets.TaskRegistry[fnName]
	return fn(args...)
}
