package main

import (
	"errors"
	"fmt"
	"math"
)

//ex1
func impostoSalario(salario, imposto float64) float64 {
	imposto /= 100
	return salario * imposto
}

//ex2
func alunosMedia(notas ...int) (float64, error) {
	resultado, quantidadeNotas := 0, 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("só é possível entrar com valores positivos.")
		}
		quantidadeNotas++
		resultado += nota
	}
	return float64(resultado / quantidadeNotas), nil
}

//ex3
const (
	C = 1000
	B = 1500
	A = 3000
)

func checarSePassouAQuantidadeDeHoras(categoria, qtdHoras float64) float64 {
	if categoria == B && qtdHoras > 160 {
		return 1.20
	}
	if categoria == A && qtdHoras > 160 {
		return 1.50
	}
	return 1
}
func calcularSalarioFuncionario(categoria, qtdHoras float64) float64 {
	switch categoria {
	case C:
		return categoria * qtdHoras
	case B:
		return (categoria * qtdHoras) * checarSePassouAQuantidadeDeHoras(categoria, qtdHoras)
	case A:

		return (categoria * qtdHoras) * checarSePassouAQuantidadeDeHoras(categoria, qtdHoras)
	}
	return 0
}

//ex 4
const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func opMinimum(x, y float64) float64 {
	return math.Min(x, y)
}
func opMaximum(x, y float64) float64 {
	return math.Max(x, y)
}

func opAverage(x, y float64) float64 {

}

func operation(operacao string) func float64 {
	switch operacao {
	case minimum:
		return opMinimum
	case maximum:
		return valores, opMaximum
	}
	return 0
}

func retornoDaOperacao(valores []float64, operacao func(value1, value2 float64) float64) float64 {
	var resultado float64
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacao(resultado, valor)
		}
	}
	return resultado
}

func main() {
	fmt.Println("----------------Exercício-01--------------")
	const (
		salarioFunc1  = 50000
		descontoFunc1 = 17
		salarioFunc2  = 150000
		descontoFunc2 = 10
	)
	var employees = map[float64]float64{salarioFunc1: descontoFunc1, salarioFunc2: descontoFunc2}
	for salario, desconto := range employees {
		fmt.Println("Salário do funcionário R$", salario, "\nDesconto", desconto, "%")
		fmt.Println("Imposto do salário de R$", impostoSalario(salario, desconto), "\n")
	}
	fmt.Println("------------------------------------------")

	fmt.Println("----------------Exercício-02--------------")
	alunoEduardo, err := alunosMedia(8, 9, -1, 7, 5)
	if err != nil {
		fmt.Println("Operação deu erro,", err)
	} else {
		fmt.Println("A média do aluno é", alunoEduardo)
	}

	fmt.Println("------------------------------------------")

	fmt.Println("----------------Exercício-03--------------")
	fmt.Println(calcularSalarioFuncionario(C, 162))
	fmt.Println(calcularSalarioFuncionario(B, 176))
	fmt.Println(calcularSalarioFuncionario(A, 172))

	fmt.Println("------------------------------------------")

	fmt.Println("----------------Exercício-03--------------")

	minhaFunc := operation(minimum)
	averageFunc := operation(average)
	maxFunc := operation(maximum)

	if err != nil {
		fmt.Println("Error:", err)
	}

	minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue)

	fmt.Println("------------------------------")
}
