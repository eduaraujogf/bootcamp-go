package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func insertionSort(items []int, out chan<- time.Duration) {
	start := time.Now()
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	finish := time.Now()
	fmt.Println("End of the insertion sort:", finish.Sub(start))
	out <- finish.Sub(start)

}

func selectionSort(items []int, out chan<- time.Duration) {
	start := time.Now()
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
	finish := time.Now()
	fmt.Println("End of the insertion sort:", finish.Sub(start))
	out <- finish.Sub(start)
}

func groupSort(items []int, out chan<- time.Duration) {
	start := time.Now()
	sort.Ints(items)
	finish := time.Now()
	fmt.Println(group)
}

func main() {
	hundredNumbers := rand.Perm(100)
	thousandNumbers := rand.Perm(1000)
	tenThousandsNumbers := rand.Perm(10000)
}
