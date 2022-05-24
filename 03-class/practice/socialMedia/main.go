package main

import "fmt"

type user struct {
	name     string
	surname  string
	age      uint
	email    string
	password string
}
type users []*user

func (u *user) changeName(name, surname string) {
	u.name = name
	u.surname = surname
}

func (u *user) changeAge(age uint) {
	u.age = age
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func (u *user) changePassword(password string) {
	u.password = password
}

func main() {
	pedrinho := user{"Pedrinho", "Silva", 30, "pedrinho123@hotmail", "123456"}

	fmt.Println(pedrinho)
	pedrinho.changeName("Joãozinho", "Filho")
	pedrinho.changeAge(31)
	pedrinho.changeEmail("pedrinho564@gmail")
	pedrinho.changePassword("912131")
	fmt.Println(pedrinho)

}
