package main

import "fmt"

var myVariable interface{}

type HeterogeneousList struct {
	Data []interface{}
}

func main() {
	l := HeterogeneousList{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "Olá")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data)
}
