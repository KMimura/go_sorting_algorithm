package main

import "fmt"

func main() {
	testVal := []int{5, 4, 7, 3, 0, 9, 2, 1, 8, 6}
	resultVal := bubbleSort(testVal)
	fmt.Println(resultVal)
}

func bubbleSort(slice []int) []int {
	for i := len(slice) - 1; i >= 1; i-- {
		if slice[i-1] > slice[i] {
			slice[i-1], slice[i] = slice[i], slice[i-1]
		}
	}
	if len(slice) > 0 {
		slice = append(slice[:1], bubbleSort(slice[1:])...)
	}
	return slice
}
