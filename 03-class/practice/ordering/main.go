package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	insertion = "insertion"
	selection = "selection"
	group     = "group"
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
	fmt.Println(len(items), "- insertion sort -", finish.Sub(start))
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
	fmt.Println(len(items), "- selection sort -", finish.Sub(start))
	out <- finish.Sub(start)
}

func groupSort(items []int, out chan<- time.Duration) {

	start := time.Now()
	sort.Ints(items)
	finish := time.Now()
	fmt.Println(len(items), "- group sort -", finish.Sub(start))
	out <- finish.Sub(start)
}

func order(order string, hundred []int, thousand []int, tenThousands []int, out chan<- time.Duration) {
	switch order {
	case insertion:
		defer insertionSort(hundred, out)
		defer insertionSort(thousand, out)
		defer insertionSort(tenThousands, out)
	case group:
		defer groupSort(hundred, out)
		defer groupSort(thousand, out)
		defer groupSort(tenThousands, out)
	case selection:
		defer selectionSort(hundred, out)
		defer selectionSort(thousand, out)
		defer selectionSort(tenThousands, out)
	default:
		panic("invalid operation")
	}
}

func main() {

	hundred := rand.Perm(100)
	thousand := rand.Perm(1000)
	tenThousands := rand.Perm(10000)
	insertionCh := make(chan time.Duration)
	groupCh := make(chan time.Duration)
	selectionCh := make(chan time.Duration)

	go order(insertion, hundred, thousand, tenThousands, insertionCh)
	<-insertionCh
	<-insertionCh

	go order(group, hundred, thousand, tenThousands, groupCh)
	<-groupCh
	<-groupCh

	go order(selection, hundred, thousand, tenThousands, selectionCh)
	<-selectionCh
	<-selectionCh

}
