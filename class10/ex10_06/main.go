package main

import (
	"errors"
	"fmt"
)

// MyError ...
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func main() {
	myErr := &MyError{
		Code:    100,
		Message: "error message",
	}
	var ptrMyErr *MyError

	err := fmt.Errorf("test error: %w", myErr)
	fmt.Println("err is myErr:", errors.Is(err, myErr))
	fmt.Println("err as MyError type:", errors.As(err, &ptrMyErr))

	otherErr := fmt.Errorf("other error: %w", errors.New("another error"))
	fmt.Println("otherErr is myErr:", errors.Is(otherErr, myErr))
	fmt.Println("otherErr as MyError type:", errors.As(otherErr, &ptrMyErr))

	testErr := errors.Unwrap(err)

	if testErr == nil {
		fmt.Println("no internal error")
	} else {
		fmt.Println(testErr.Error())
	}

	testErr = errors.Unwrap(otherErr)
	if testErr == nil {
		fmt.Println("no internal error")
	} else {
		fmt.Println(testErr.Error())
	}

	testErr = errors.Unwrap(errors.New("error"))
	if testErr == nil {
		fmt.Println("no internal error")
	} else {
		fmt.Println(testErr.Error())
	}

}
