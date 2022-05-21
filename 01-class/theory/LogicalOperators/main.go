package main

import "fmt"

func main() {
	x, y, z := 10, 20, 30

	fmt.Println("x < y && x > z =", x < y && x > z)
	fmt.Println("x < y || x > z =", x < y || x > z)
	fmt.Println("!(x == y && x > z) =", !(x == y && x > z))

}
