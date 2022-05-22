package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type myStructure struct {
	Field1 string `myLabel:"value"`
	Field2 string `myLabel:"value"`
	Field3 string `myLabel:"value"`
}

type Person struct {
	FirstName string `json:"first_name"`
	Surname   string `json:"surname"`
	Telephone string `json:"telephone"`
	Address   string `json:"address"`
}

type PersonalizedPerson struct {
	FirstName string `bd:"first_name"`
	Surname   string `bd:"surname"`
	Telephone string `bd:"telephone"`
	Address   string `bd:"address"`
}

func main() {
	p := Person{"Paula", "Monteiro", "43434343", "Street Limoeiro 123"}
	myJson, err := json.Marshal(p)
	fmt.Println(string(myJson))
	fmt.Println(err)

	person := PersonalizedPerson{}
	pl := reflect.TypeOf(person)
	fmt.Println("Type:", pl.Name())
	fmt.Println("Kind:", pl.Kind())

	for i := 0; i < pl.NumField(); i++ {
		field := pl.Field(i)
		tag := field.Tag.Get("bd")
		fmt.Println(tag)
	}
}
