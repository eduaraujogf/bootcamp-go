package pkg

import (
	"fmt"
)

type Student struct {
	Name          string
	Surname       string
	RG            string
	AdmissionDate string
}

func (s Student) Details() {
	fmt.Printf("Name: %s\nSurname: %s\nRG: %s\nAdmission Date: %s\n\n", s.Name, s.Surname, s.RG, s.AdmissionDate)
}
