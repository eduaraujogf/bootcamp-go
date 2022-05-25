package main

import (
	"errors"
	"fmt"
	"log"
)

func sayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("error: empty name.")
	}
	return "Hello, " + name, nil
}

func main() {
	name, err := sayHello("")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(name)
}
