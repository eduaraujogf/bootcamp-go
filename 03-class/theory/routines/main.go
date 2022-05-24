package main

import (
	"fmt"
	"time"
)

func process(i int) {
	fmt.Println(i, "-Start")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-End")
}

func main() {
	// process(10)

	for i := 0; i < 10; i++ {
		go process(i)
	}
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("End of the program")
}
