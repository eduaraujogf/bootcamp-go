package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func redis() {
	fmt.Println("closing redis connection...")
}
func mysql() {
	fmt.Println("closing mysql connection...")
}
func mongo() {
	fmt.Println("closing mongo connection...")
}

func main() {
	defer hello()
	fmt.Println("after")

	defer redis()
	defer mysql()
	defer mongo()
}
