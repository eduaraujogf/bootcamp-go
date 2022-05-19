package main

import "fmt"

func main() {

	fmt.Println("---------------Exercício 1 -------------------")
	palavra := "carro"

	fmt.Println("Número de letras:", len(palavra))
	for _, letra := range palavra {
		fmt.Println(string(letra))
	}

	fmt.Println("----------------------------------------------")

	fmt.Println("---------------Exercício 2 -------------------")
	idade := 23
	maisDeUmAnoDeAtividade := true
	salario := 90000

	switch {
	case (idade > 22 && maisDeUmAnoDeAtividade) && (salario > 100000):
		fmt.Println("Você está elegível para o empréstimo e não precisará pagar juros por ele.")
	case idade > 22 && maisDeUmAnoDeAtividade:
		fmt.Println("Você está elegível para o empréstimo, porém terá que pagar juros por ele.")
	default:
		fmt.Println("Você não está elegível para o empréstimo.")
	}

	fmt.Println("----------------------------------------------")
	fmt.Println("---------------Exercício 3 -------------------")
	// var mes uint8
	meses := map[uint8]string{
		0:  "Janeiro",
		1:  "Fevereiro",
		2:  "Março",
		3:  "Abril",
		4:  "Maio",
		5:  "Junho",
		6:  "Julho",
		7:  "Agosto",
		8:  "Setembro",
		9:  "Outubro",
		10: "Novembro",
		11: "Dezembro",
	}
	// fmt.Scan(&mes)
	fmt.Println("O número", 1, "corresponde ao mês de", meses[1])

	fmt.Println("----------------------------------------------")

	fmt.Println("---------------Exercício 4 -------------------")

	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	var counter int

	fmt.Println("O Benjamin tem", employees["Benjamin"], "anos.")

	for _, employeeAge := range employees {
		if employeeAge > 21 {
			counter++
		}
	}
	fmt.Println("Existem", counter, "funcionários com mais de 21 anos.")
	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)

	fmt.Println("----------------------------------------------")
}
