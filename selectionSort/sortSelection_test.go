package main

import (
	"testing"
)

func TestSelectionSort1(t *testing.T) {
	testCase := []int{8, 5, 6, 4, 1, 9, 0, 2, 3, 7}
	sortChan := make(chan []int)
	go selectionSort(testCase, sortChan)
	result := <-sortChan
	testFailureFlg := false
	expectedResult := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, e := range expectedResult {
		if result[i] != e {
			testFailureFlg = true
		}
	}
	if testFailureFlg {
		t.Fatalf("test1 failed")
	}
}

func TestSelectionSort2(t *testing.T) {
	testCase := []int{8, 5, 5, 6, 4, 1, 9, 0, 2, 3, 4, 7}
	sortChan := make(chan []int)
	go selectionSort(testCase, sortChan)
	result := <-sortChan
	testFailureFlg := false
	expectedResult := []int{0, 1, 2, 3, 4, 4, 5, 5, 6, 7, 8, 9}
	for i, e := range expectedResult {
		if result[i] != e {
			testFailureFlg = true
		}
	}
	if testFailureFlg {
		t.Fatalf("test2 failed")
	}
}

func TestSelectionSort3(t *testing.T) {
	testCase := []int{0}
	sortChan := make(chan []int)
	go selectionSort(testCase, sortChan)
	result := <-sortChan
	testFailureFlg := false
	expectedResult := []int{0}
	for i, e := range expectedResult {
		if result[i] != e {
			testFailureFlg = true
		}
	}
	if testFailureFlg {
		t.Fatalf("test3 failed")
	}
}
