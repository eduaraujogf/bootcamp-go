package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	err := os.Setenv("Name", "Gopher")
	if err != nil {
		log.Println("error to set environmental variable.", err)
	}
	value := os.Getenv("Name")

	log.Println(value)

	valueLook, ok := os.LookupEnv("NAME")

	if ok {
		log.Println(ok)
	} else {
		log.Println(valueLook)
	}

	files, err := os.ReadDir(".")
	if err != nil {
		log.Println("directory error:", err)
	}

	for _, file := range files {
		log.Printf("%s is directory? %t", file.Name(), file.Type().IsDir())
	}

	data, err := os.ReadFile("./README.md")
	if err != nil {
		log.Println("error to read file:", err)
	}
	log.Println(string(data))
	log.Printf("%s", data)

}
