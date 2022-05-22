package main

import "fmt"

type Preferences struct {
	Foods  string
	Movie  string
	Series string
	Anime  string
	Sports string
}

type Person struct {
	Name       string
	Gender     string
	Age        int
	Profession string
	Weight     float64
	Tastes     Preferences
}

func main() {
	p1 := Person{"Paula", "Female", 34, "Engineer", 65.5, Preferences{"Chicken", "Titanic", "", "", ""}}
	p2 := Person{
		Name:       "Vitor",
		Gender:     "Male",
		Age:        30,
		Profession: "Architect",
		Weight:     77,
		Tastes: Preferences{
			Foods: "Roasted Chicken",
			Movie: "Fight Club",
			Anime: "Shingeki no kyojin",
		},
	}
	p2.Weight = 70
	p2.Tastes.Sports = "Football"
	var p3 Person
	p3.Name = "Fernando"
	p3.Age = 15
	p3.Tastes = Preferences{Foods: "Vegetables", Movie: "Ice age"}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

}
