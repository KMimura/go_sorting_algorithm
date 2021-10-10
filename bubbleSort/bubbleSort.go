package main

import "fmt"

func main() {
	testVal := []int{5, 4, 7, 3, 0, 9, 2, 1, 8, 6}
	resultChan := make(chan []int)
	go bubbleSort(testVal, resultChan)
	resultSlice := <-resultChan
	fmt.Println(resultSlice)
}

func bubbleSort(slice []int, c chan []int) {
	for i := len(slice) - 1; i >= 1; i-- {
		if slice[i-1] > slice[i] {
			slice[i-1], slice[i] = slice[i], slice[i-1]
		}
	}
	if len(slice) > 0 {
		sortChan := make(chan []int)
		go bubbleSort(slice[1:], sortChan)
		sortedSlice := <-sortChan
		slice = append(slice[:1], sortedSlice...)
	}
	c <- slice
	close(c)
}
