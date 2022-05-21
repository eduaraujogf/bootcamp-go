package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var s = []bool{true, false}

	fmt.Println(s[0])

	b := make([]int, 5) // len (a) = 5
	fmt.Println(b)

	primes := []int{2, 3, 5, 7, 11, 13}
	// If we do not put a value after the : it will go until the end of the slice
	fmt.Println(primes[1:])
	fmt.Println(primes[2:3])
	fmt.Println(primes[:3])

	var j []int

	j = append(j, 2, 3, 5)
}
