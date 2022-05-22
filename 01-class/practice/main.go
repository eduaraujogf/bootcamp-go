package main

import (
	"fmt"

	"github.com/eduaraujogf/bootcamp-go/01-class/practice/pkg"
)

func main() {
	//Morning
	fmt.Println("----------------Morning--------------\n")
	fmt.Println("----------------Exercise-01--------------")
	pkg.PrintName()
	fmt.Println("------------------------------------------")

	fmt.Println("----------------Exercise-02--------------")
	pkg.Climate()
	fmt.Println("------------------------------------------")

	fmt.Println("----------------Exercise-03--------------")
	// var nome string
	// var sobrenome string
	// var idade int
	// sobrenome = "Sobrenome"
	// var licencaParaDirigir = true
	// var estaturaDaPessoa int
	// quantidadeDeFilhos := 2
	fmt.Println("------------------------------------------")

	fmt.Println("----------------Exercise-04--------------\n")
	// var sobrenome string = "Silva"
	// var idade uint8 = 25
	// boolean := false
	// var salario float64 = 4585.90
	// var nome string = "Fellipe"

	//Afternoon
	fmt.Println("----------------Afternoon--------------\n")
	fmt.Println("----------------Exercise-01--------------")
	pkg.LettersOfAWord("Car")
	fmt.Println("------------------------------------------")
	fmt.Println("----------------Exercise-02--------------")
	pkg.Loan(23, true, 110000)
	fmt.Println("------------------------------------------")
	fmt.Println("----------------Exercise-03--------------")
	pkg.PrintMonth(8)
	fmt.Println("------------------------------------------")
	fmt.Println("----------------Exercise-04--------------")
	pkg.HowOldIs()
	fmt.Println("------------------------------------------")

}
