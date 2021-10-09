package main

import "fmt"

func main() {
	testVal := []int{5, 1, 0, 8, 2, 7, 4, 3, 9, 6}
	fmt.Println(quickSort(testVal))
}

func quickSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	rootVal := slice[0]
	var ltRoot []int
	var gtRoot []int
	for _, a := range slice {
		if a < rootVal {
			ltRoot = append(ltRoot, a)
		}
		if a > rootVal {
			gtRoot = append(gtRoot, a)
		}
	}
	var ltRootSorted []int
	var gtRootSorted []int

	ltRootSorted = quickSort(ltRoot)
	gtRootSorted = quickSort(gtRoot)
	ltRootSorted = append(ltRootSorted, rootVal)
	returnVal := append(ltRootSorted, gtRootSorted...)
	return returnVal
}
