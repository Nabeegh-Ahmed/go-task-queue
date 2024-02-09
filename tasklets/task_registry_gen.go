package tasklets

type AnyFunc func(...interface{}) (interface{}, error)

var TaskRegistry = map[string]AnyFunc{
	"IsPrime": IsPrime,
}
