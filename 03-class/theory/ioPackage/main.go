package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	d1 := []byte("\n\thello, gophers!")
	// errWrite := os.WriteFile("./myFile.txt", d1, 0644)

	// if errWrite != nil {
	// 	log.Println("Error to write in the file:", errWrite)
	// }
	f, errW := os.OpenFile("./myFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if errW != nil {
		panic(errW)
	}

	defer f.Close()

	if _, err := f.WriteString(string(d1)); err != nil {
		panic(err)
	}

	rs := strings.NewReader("some io.Reader stream to be read\n")
	_, errc := io.Copy(os.Stdout, rs)

	if errc != nil {
		log.Println(errc)
	}

	_, errS := io.WriteString(os.Stdout, "Hello world!")
	if errS != nil {
		log.Println(errS)
	}
}
