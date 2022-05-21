package pkg

import "fmt"

func LettersOfAWord(word string) {
	fmt.Println("Numbers of letters:", len(word))
	for _, letter := range word {
		fmt.Println(string(letter))
	}
}
