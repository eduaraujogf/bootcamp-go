package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("error number 1")
var err2 = errors.New("error number 2")

func x() error {
	return fmt.Errorf("extra information about the error: %w", err1)
}

func main() {
	e := x()
	coincidence := errors.Is(e, err1)
	fmt.Println(coincidence)

}
