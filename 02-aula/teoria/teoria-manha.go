package main

import (
	"errors"
	"fmt"
)

// func soma(a, b float32) float32 {
// 	return a + b

// }

const (
	Soma   = "+"
	Subtra = "-"
	Multip = "*"
	Divis  = "/"
)

func operacaoAitmetica(valor1, valor2 float64, operador string) float64 {
	switch operador {
	case Soma:
		return valor1 + valor2
	case Subtra:
		return valor1 - valor2
	case Multip:
		return valor1 * valor2
	case Divis:
		if valor2 != 0 {
			return valor1 / valor2
		}
	}
	return 0
}

// func soma(values ...float64) float64 {
// 	var resultado float64
// 	for _, value := range values {
// 		resultado += value
// 	}
// 	return resultado
// }

func helloWorld(value string) string {
	return "Hello World - " + value
}

func helloGo(value string) string {
	return "Hello Go - " + value
}

func foo(str string, fn func(value string) string) string {
	return fn(str)
}

func division(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("o divisor não pode ser zero")
	}
	return dividendo / divisor, nil
}

func operacoes(valor1, valor2 float64) (soma, subtra, multip, divis float64) {
	soma = valor1 + valor2
	subtra = valor1 - valor2
	multip = valor1 * valor2

	if valor2 != 0 {
		divis = valor1 / valor2
	}

	return soma, subtra, multip, divis

}

func main() {
	// fmt.Println(soma(5, 5))
	// fmt.Println(operacaoAitmetica(32, 5, Soma))
	// fmt.Println(operacaoAitmetica(10, 5, Subtra))
	// fmt.Println(operacaoAitmetica(5, 5, Multip))
	// fmt.Println(operacaoAitmetica(5, 5, Divis))
	// // fmt.Println(soma(100, 200, 300, 400))
	// fmt.Println(foo("Eduardo", helloWorld))
	// fmt.Println(foo("Marmota", helloGo))

	// res, err := division(2, 0)

	// if err != nil {
	// 	fmt.Println("Operação deu erro:", err.Error())
	// } else {
	// 	fmt.Println("operação deu certo:", res)
	// }

	soma, sub, mult, div := operacoes(10, 5)

	fmt.Println(soma)
	fmt.Println(sub)
	fmt.Println(mult)
	fmt.Println(div)

}
