package main

import "fmt"

func Increase(v *int) {
	*v++
}

func main() {
	var p1 *int
	var p2 = new(int)

	fmt.Println("The address of the variable p1 is:", p1)
	fmt.Println("The address of the variable p2 is:", p2)

	// * To create a variable of type pointer
	// * To visualize the position in memory
	// & To set a value to the pointer

	var v int = 19
	var p *int = &v

	fmt.Printf("The pointer p  memory value is: %v\n", p)
	fmt.Printf("The pointer p value is: %d\n", *p)

	var j int = 19

	Increase(&j)
	fmt.Println("The value of j is: ", j)

}
