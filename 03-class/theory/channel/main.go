package main

import (
	"fmt"
	"net/http"
	"time"
)

func process(i int, c chan int) {
	fmt.Println(i, "-Start")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-End")
	c <- i
}

func checkStatus(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error to access the url: %s", url)
		ch <- fmt.Sprintf("Error: %s", err)
		return
	}
	defer resp.Body.Close()

	u := fmt.Sprintf("url: %s - status: %s", url, resp.Status)
	ch <- u
}

func sum(grades []int, out chan<- int) {
	var total int

	for _, grade := range grades {
		total += grade
	}

	out <- total
	close(out)

}

func avg(in <-chan int, out chan<- int, length int) {
	var avgGrades int

	for v := range in {
		avgGrades = v / length
	}
	out <- avgGrades

	close(out)
}

func printAvg(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	grades := []int{10, 5, 6, 3, 7, 8}
	sumCh := make(chan int)
	avgCh := make(chan int)

	go sum(grades, sumCh)
	go avg(sumCh, avgCh, len(grades))
	printAvg(avgCh)

	// urls := []string{
	// 	"https://www.spotify.com",
	// 	"https://www.apple.com",
	// 	"https://www.google.com",
	// 	"https://www.microsoft.com",
	// 	"https://globoesporte.globo.com",
	// }
	// c := make(chan int)
	// ch := make(chan string)

	// for _, url := range urls {
	// 	go checkStatus(url, ch)
	// }

	// for range urls {
	// 	fmt.Println(<-ch)
	// }

	// for i := 0; i < 10; i++ {
	// 	go process(i, c)
	// }

	// variable := <-c
	// fmt.Println("End of the program", variable)
}
