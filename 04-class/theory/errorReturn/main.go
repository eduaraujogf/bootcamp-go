package main

import (
	"fmt"
	"net/http"
)

type errorURL struct {
	msg        string
	statusCode int
}

func (e *errorURL) Error() string {
	return fmt.Sprintf("customized error. msg: %s, code: %d", e.msg, e.statusCode)
}

func checkURL(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, fmt.Errorf("it occurred an error to access the URL: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, &errorURL{msg: "status different from 200", statusCode: resp.StatusCode}
	}
	return true, nil
}

func main() {
	result, err := checkURL("https://www.google.com.br")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Status:", result)
}
