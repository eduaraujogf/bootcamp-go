package main

import (
	"fmt"
	"net/http"
	"os"
)

type myError struct {
	status int
	msg    string
}

func (m *myError) Error() string {
	return fmt.Sprintf("%d | %v", m.status, m.msg)
}

func myErrorTest(status int) (int, error) {
	if status >= http.StatusBadRequest {
		return http.StatusInternalServerError, &myError{
			status: status,
			msg:    "Something went wrong",
		}
	}
	return http.StatusOK, nil
}

func main() {

	status, err := myErrorTest(400)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Status %d, it works!\n", status)

}
