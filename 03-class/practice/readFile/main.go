package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

func readCSV(filepath string) {
	file, err := os.Open("./03-class/practice/fileStore/" + filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := tabwriter.NewWriter(os.Stdout, 20, 30, 1, '\t', tabwriter.AlignRight)

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	header := strings.Split(scanner.Text(), ",")

	for _, c := range header {
		fmt.Fprintf(w, "%s\t\t", c)
	}

	fmt.Fprintln(w)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for _, v := range values {
			fmt.Fprintf(w, "%s\t\t", v)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
}

func main() {
	readCSV("products.csv")

}
