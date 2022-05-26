package main

import "fmt"

func debugger() {
	err := recover()

	if err != nil {
		fmt.Println(err)
	}
}

func isEven(num int) {
	defer debugger()

	if (num % 2) != 0 {
		panic("is not even")
	}
	fmt.Println(num, "is a number even!")
}

func main() {
	isEven(1)

	fmt.Println("End of the program")
}
