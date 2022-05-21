package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	fruits := []string{"apple", "banana", "avocado"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// for {
	// 	sum++
	// }

	sumWhile := 1

	for sumWhile < 10 {
		sumWhile += sumWhile
		fmt.Println(sumWhile)
	}
	fmt.Println(sumWhile)

	sumBreak := 0
	for {
		sumBreak++
		if sumBreak >= 1000 {
			break
		}
	}
	fmt.Println(sumBreak)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "is odd")
	}

}
