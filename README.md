# Task Execution Service
A task execution service written in Golang. It consumes a RabbitMQ queue and executes tasks inside a worker. The service provides an easy to use interface to pass function pointers and arguments for tasks to be executed inside the service. 

# Get Started
- Clone the repo
- Add your tasks in `tasklets`
- Add a wrapper for your tasklet that takes in arbitrary arguments. Use the following template as example
  ```go
  func SomeFuncWrapper(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("SomeFunc expects 1 arguments")
	}
  // Check for types
	a, ok1 := args[0].(int32)
	if !ok1 {
		return nil, fmt.Errorf("invalid argument types for SomeFunc")
	}
  // Execute the main function
	return someFunc(a), nil
  ```
- Run `go generate`
- Now you can enqueue this task as you please by using
  ```go
  execution := taskExecution.TaskExecutionInit()
	execution.Enqueue(tasklets.IsPrime, 5)
  ```
- You can spawn more workers to execute tasks using `go worker.Worker()`
