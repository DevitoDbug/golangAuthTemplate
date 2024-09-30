package utils

import "fmt"

type ErrorContext struct {
	Context string 
	Value string 
}

func (e *ErrorContext) Error() string {
    return fmt.Sprintf("Context: %s, Value: %s", e.Context, e.Value)
}