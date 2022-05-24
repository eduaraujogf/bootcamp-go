package main

import "fmt"

/*
	%v  default
	%T 	type
	%t	bool
	%s	string
	%f	float
	%d 	int
	%b	binary
	%o	octal
	%c	char
	%p	memory

*/

func main() {
	name, age := "Kim", 22
	fmt.Print(name, " is ", age, " years old.\n")
	fmt.Println(name, "is", age, "years old.")
	fmt.Printf("%10.2f\n", 12222.123123)
	fmt.Printf("%.2f\n", 12222.123123)
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%10d\n", 12222)
	fmt.Printf("%10s\n", "aa")

	res := fmt.Sprint(name, " is ", age, " years old.\n")
	fmt.Print(res)

	res = fmt.Sprintf("%s is %d years old.", name, age)
	fmt.Println(res)

}
