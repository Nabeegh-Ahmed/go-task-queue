package tasklets

import "fmt"

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// go:tasklet
func IsPrime(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("IsPrime expects 2 arguments")
	}
	a, ok1 := args[0].(int)
	if !ok1 {
		return nil, fmt.Errorf("invalid argument types for exampleFunc")
	}
	return isPrime(a), nil
}
