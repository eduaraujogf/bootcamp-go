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

func totalCalculator(priceString string, quantityString string) float64 {
	price, _ := strconv.ParseFloat(strings.TrimSpace(priceString), 64)
	quantity, _ := strconv.Atoi(strings.TrimSpace(quantityString))
	return price * float64(quantity)
}

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
				total += totalCalculator(values[1], values[2])
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
