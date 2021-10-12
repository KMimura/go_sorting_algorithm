package main

import "fmt"

func main() {
	testVal := []int{5, 4, 7, 3, 0, 9, 2, 1, 8, 6}
	sortChan := make(chan []int)
	go selectionSort(testVal, sortChan)
	resultVal := <-sortChan
	fmt.Println(resultVal)
}

func selectionSort(slice []int, c chan []int) {
	if len(slice) <= 1 {
		c <- slice
		close(c)
		return
	}
	minimumVal := slice[0]
	var minimumValIndex int
	for i, s := range slice {
		if s < minimumVal {
			minimumVal = s
			minimumValIndex = i
		}
	}
	if minimumValIndex != 0 {
		slice[0], slice[minimumValIndex] = slice[minimumValIndex], slice[0]
	}
	resultChan := make(chan []int)
	go selectionSort(slice[1:], resultChan)
	sortedSlice := <-resultChan
	c <- append(slice[:1], sortedSlice...)
}
