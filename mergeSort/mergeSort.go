package main

import "fmt"

func main() {
	testVal := []int{5, 4, 7, 3, 0, 9, 2, 1, 8, 6}
	resultChan := make(chan []int)
	go mergeSort(testVal, resultChan)
	resultSlice := <-resultChan
	fmt.Println(resultSlice)
}

func mergeSort(slice []int, c chan []int) {
	sliceLen := len(slice)
	if sliceLen <= 1 {
		c <- slice
		close(c)
		return
	}
	firstChan := make(chan []int)
	secondChan := make(chan []int)
	go mergeSort(slice[:sliceLen/2], firstChan)
	go mergeSort(slice[sliceLen/2:], secondChan)
	firstSliceSorted := <-firstChan
	secondSliceSorted := <-secondChan
	var concatenatedSlice []int
	firstSliceCursor := 0
	secondSliceCursor := 0
	for firstSliceCursor < len(firstSliceSorted) || secondSliceCursor < len(secondSliceSorted) {
		if firstSliceCursor == len(firstSliceSorted) {
			concatenatedSlice = append(concatenatedSlice, secondSliceSorted[secondSliceCursor:]...)
			break
		} else if secondSliceCursor == len(secondSliceSorted) {
			concatenatedSlice = append(concatenatedSlice, firstSliceSorted[firstSliceCursor:]...)
			break
		}
		if firstSliceSorted[firstSliceCursor] < secondSliceSorted[secondSliceCursor] {
			concatenatedSlice = append(concatenatedSlice, firstSliceSorted[firstSliceCursor])
			firstSliceCursor++
		} else if firstSliceSorted[firstSliceCursor] > secondSliceSorted[secondSliceCursor] {
			concatenatedSlice = append(concatenatedSlice, secondSliceSorted[secondSliceCursor])
			secondSliceCursor++
		} else if firstSliceSorted[firstSliceCursor] == secondSliceSorted[secondSliceCursor] {
			concatenatedSlice = append(concatenatedSlice, firstSliceSorted[firstSliceCursor])
			concatenatedSlice = append(concatenatedSlice, secondSliceSorted[secondSliceCursor])
			firstSliceCursor++
			secondSliceCursor++
		}
	}
	c <- concatenatedSlice
	close(c)
}
