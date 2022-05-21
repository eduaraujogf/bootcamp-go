package main

import "fmt"

func main() {
	myMap := map[string]int{}
	//the function make help us to initialize the map, but we will not be able to fill it with values in the same instruction
	myMapMake := make(map[string]string)

	fmt.Println(len(myMap), myMapMake)

	var students = map[string]int{"João": 20, "Pedro": 26}
	fmt.Println(students["João"])
	students["Brenda"] = 19
	students["Marcos"] = 22
	fmt.Println(students)
	delete(students, "João")
	fmt.Println(students)

	for key, element := range students {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}
