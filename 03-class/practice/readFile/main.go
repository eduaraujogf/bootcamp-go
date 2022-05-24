package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var total float64

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for i, v := range values {
			fmt.Fprintf(w, "%s\t\t", v)
			if i == 0 {
				price, _ := strconv.ParseFloat(strings.TrimSpace(values[1]), 64)
				quantity, _ := strconv.Atoi(strings.TrimSpace(values[2]))
				total += price * float64(quantity)
			}

		}

		fmt.Fprintln(w)
	}
	fmt.Fprintf(w, "\t\t%.2f\n", total)

	w.Flush()
}

func main() {
	readCSV("products.csv")

}
