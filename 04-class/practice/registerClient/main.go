package main

import (
	"errors"
	"fmt"
	"os"
)

type client struct {
	file    string
	name    string
	surname string
	rg      string
	phone   string
	address string
}

func debugger() {
	err := recover()

	if err != nil {
		fmt.Println(err)
	}
}

func (c *client) generateId(id string) {
	if id == "" {
		panic("ID could not be empty")
	}
	c.file = id
}

func (c *client) clientChecker(fileName string) error {
	defer debugger()
	file, err := os.Open(fileName)
	if err != nil {
		err = errors.New("the file was not found or it is damaged")
		panic(err)
	}
	defer file.Close()
	return nil
}

func main() {
	client := client{"", "Eduardo", "Araujo", "121314141", "9121214141", "Street"}
	client.generateId("10")
	fmt.Println(client)

	if err := client.clientChecker(client.file); err != nil {
		fmt.Println(err)
	}

}
