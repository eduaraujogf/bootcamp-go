package main

import (
	circ "bootcamp-go/02-aula/teoria/pkg/circulo"
	"fmt"
)

// type Preferencias struct {
// 	Comidas  string
// 	Filme    string
// 	Series   string
// 	Animes   string
// 	Esportes string
// }

// type Pessoa struct {
// 	Nome      string
// 	Gênero    string
// 	Idade     int
// 	Profissão string
// 	Peso      float64
// 	Gostos    Preferencias
// }

// type Pessoa struct {
// 	PrimeiroNome string `json:"primeiro_nome", bson:"primeiroNome"`
// 	Sobrenome    string `json:"sobrenome"`
// 	Telefone     string `json:"sobrenome"`
// 	Endereco     string `json:"endereco"`
// }

func main() {
	// eduardo := Pessoa{
	// 	Nome:      "Eduardo",
	// 	Gênero:    "Masculino",
	// 	Idade:     25,
	// 	Profissão: "Professor de Go",
	// 	Peso:      99,
	// 	Gostos: Preferencias{
	// 		Comidas: "Frango",
	// 		Filme:   "Titanic",
	// 	},
	// }
	// p := Pessoa{"Paula", "Monteiro", "4512121", "Rua Liomeiro 123"}

	// meuJson, err := json.Marshal(p)

	// if err != nil {
	// 	fmt.Println("Não foi possível gerar o JSON. erro:", err)
	// }

	// fmt.Println(string(meuJson))

	circulo := circ.Circulo{}
	circulo.SetarRaio(3)

	fmt.Println(circulo.GetRaio())
}
