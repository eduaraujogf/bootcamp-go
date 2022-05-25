package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
	x   string
}

func (e *myError) Error() string {
	return fmt.Sprintf("it occurred an error: %s, %s", e.msg, e.x)
}

func main() {
	e := &myError{"New error", "404"}

	var err *myError

	isMyError := errors.As(e, &err)

	fmt.Println(isMyError)

}
