package main

import "fmt"

func main() {
	//Exercício 1
	fmt.Println("----------------Exercício-01--------------")
	nome, idade := "Eduardo", 31
	fmt.Println(nome, idade)
	fmt.Println("------------------------------------------")

	//Exercício 2
	fmt.Println("----------------Exercício-02--------------")
	var (
		temperatura        float32
		umidade            float32
		pressaoAtmosferica float32
	)
	temperatura, umidade, pressaoAtmosferica = 11.5, 81.3, 1.1
	fmt.Println(temperatura, umidade, pressaoAtmosferica)
	fmt.Println("------------------------------------------")

	//Exercício 3
	fmt.Println("----------------Exercício-03--------------")
	// var nome string
	// var sobrenome string
	// var idade int
	// sobrenome = "Sobrenome"
	// var licencaParaDirigir = true
	// var estaturaDaPessoa int
	// quantidadeDeFilhos := 2
	fmt.Println("------------------------------------------")

	//Exercício 4
	fmt.Println("----------------Exercício-04--------------")
	// var sobrenome string = "Silva"
	// var idade uint8 = 25
	// boolean := false
	// var salario float64 = 4585.90
	// var nome string = "Fellipe"
	fmt.Println("------------------------------------------")
}
