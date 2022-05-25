package main

import "fmt"

func main() {
	status := 401
	message := "an error occurred"

	err := fmt.Errorf("status - %d - message: %s", status, message)
	fmt.Println(err)
}
