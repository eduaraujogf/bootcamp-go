package main

import (
	"errors"
	"fmt"
)

type errorTwo struct{}

func (e errorTwo) Error() string {
	return "error two happened"
}

func main() {
	e2 := errorTwo{}
	e1 := fmt.Errorf("e2: %w", e2)
	fmt.Println(errors.Unwrap(e1))
	fmt.Println(errors.Unwrap(e2))

}
