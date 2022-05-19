package main

import "fmt"

func main() {
	x, y := 10, 20

	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x mod y = %d\n", x%y)

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n", y)

	idade := 10

	fmt.Println(&idade) // endereço de memória

	//arrays

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//slices
	b := make([]string, 3)

	b[0] = "Hello"
	b[1] = "World"
	b[2] = "Go"
	b = append(b, "append")

	fmt.Println(b[0], b[1], b[2], b[3])
	fmt.Println(b)

	var c []string
	c = append(c, "Hello")
	c = append(c, "World")
	c = append(c, "Go")
	c = append(c, "Slice")

	fmt.Println(c[0], c[1], c[2], c[3])
	fmt.Println(c)

	//map
	studentsMake := make(map[string]string) // não passa valor defaults
	fmt.Println(len(studentsMake))

	var students = map[string]int{"João": 20, "Pedro": 26}
	fmt.Println(students["João"])
	fmt.Println(len(students))
	students["Brenda"] = 19
	fmt.Println(students)

	delete(students, "João") //só paar o map
	fmt.Println(students)

	//Percorrendo o Map
	for key, element := range students {
		fmt.Println("Chave:", key, "- Valor:", element)
	}

	//For
	frutas := []string{"maçã", "banana", "pêra"}
	for _, value := range frutas {
		fmt.Println(value)
	}

	if result := sum(3, 5); result < 10 {
		fmt.Println("Valor menor que 10")
	}

	//switch case

	dias := 8

	switch dias {
	case 0:
		fmt.Println("Segunda-feira")
	case 1:
		fmt.Println("Terça-feira")
	case 2:
		fmt.Println("Quarta-feira")
	case 3:
		fmt.Println("Quinta-feira")
	case 4:
		fmt.Println("Sexta-feira")
	case 5:
		fmt.Println("Sábado")
	case 6:
		fmt.Println("Domingo")
	default:
		fmt.Println("Erro")
	}

}

func sum(a int, b int) int {
	return a + b
}
