package main

import "fmt"

func main() {
	testVal := []int{5, 1, 0, 8, 2, 7, 4, 3, 9, 6}
	sortChan := make(chan []int)
	go quickSort(&testVal, sortChan)
	resultVal := <-sortChan
	fmt.Println(resultVal)
}

func quickSort(slice *[]int, c chan []int) {
	if len(*slice) <= 1 {
		c <- *slice
		close(c)
		return
	}
	rootVal := (*slice)[0]
	// less than
	var ltRoot []int
	// greater than
	var gtRoot []int
	// same as
	var saRoot []int
	for _, a := range *slice {
		if a < rootVal {
			ltRoot = append(ltRoot, a)
		} else if a > rootVal {
			gtRoot = append(gtRoot, a)
		} else if a == rootVal {
			saRoot = append(saRoot, a)
		}
	}

	ltChan := make(chan []int)
	gtChan := make(chan []int)

	go quickSort(&ltRoot, ltChan)
	go quickSort(&gtRoot, gtChan)

	ltRootSorted := <-ltChan
	gtRootSorted := <-gtChan

	ltRootSorted = append(ltRootSorted, saRoot...)
	returnVal := append(ltRootSorted, gtRootSorted...)
	c <- returnVal
	close(c)
}
