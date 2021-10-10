package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	testCase := []int{8, 5, 6, 4, 1, 9, 0, 2, 3, 7}
	sortChan := make(chan []int)
	go bubbleSort(testCase, sortChan)
	result := <-sortChan
	testFailureFlg := false
	expectedResult := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, e := range expectedResult {
		if result[i] != e {
			testFailureFlg = true
		}
	}
	if testFailureFlg {
		fmt.Println("test failed")
		fmt.Println(result)
	} else {
		fmt.Println("test success")
	}
}

func TestBubbleSort2(t *testing.T) {
	testCase := []int{8, 5, 5, 6, 4, 1, 9, 0, 2, 3, 4, 7}
	sortChan := make(chan []int)
	go bubbleSort(testCase, sortChan)
	result := <-sortChan
	testFailureFlg := false
	expectedResult := []int{0, 1, 2, 3, 4, 4, 5, 5, 6, 7, 8, 9}
	for i, e := range expectedResult {
		if result[i] != e {
			testFailureFlg = true
		}
	}
	if testFailureFlg {
		fmt.Println("test failed")
		fmt.Println(result)
	} else {
		fmt.Println("test success")
	}
}
